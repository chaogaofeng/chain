package app

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	wasmmodulekeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	tmlog "github.com/tendermint/tendermint/libs/log"

	"runtime/debug"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibcante "github.com/cosmos/ibc-go/v3/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
	permissionmodulekeeper "github.com/goldnet/chain/x/permission/keeper"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	evmante "github.com/tharsis/ethermint/app/ante"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC
// channel keeper.
type HandlerOptions struct {
	AccountKeeper   evmtypes.AccountKeeper
	BankKeeper      evmtypes.BankKeeper
	IBCKeeper       *ibckeeper.Keeper
	FeegrantKeeper  authante.FeegrantKeeper
	SignModeHandler authsigning.SignModeHandler
	SigGasConsumer  func(meter sdk.GasMeter, sig signing.SignatureV2, params authtypes.Params) error

	PermissionKeeper  permissionmodulekeeper.Keeper
	WasmConfig        *wasmtypes.WasmConfig
	TXCounterStoreKey sdk.StoreKey
	FeeMarketKeeper   evmtypes.FeeMarketKeeper
	EvmKeeper         evmante.EVMKeeper
	MaxTxGasWanted    uint64
}

// NewAnteHandler returns an ante handler responsible for attempting to route an
// Ethereum or SDK transaction to an internal ante handler for performing
// transaction-level processing (e.g. fee payment, signature verification) before
// being passed onto it's respective handler.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	//if options.AccountKeeper == nil {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for AnteHandler")
	//}
	//if options.BankKeeper == nil {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for AnteHandler")
	//}
	if options.SignModeHandler == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	if options.WasmConfig == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "wasm config is required for ante builder")
	}
	if options.TXCounterStoreKey == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "tx counter key is required for ante builder")
	}

	var sigGasConsumer = options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = evmante.DefaultSigVerificationGasConsumer
	}

	return func(
		ctx sdk.Context, tx sdk.Tx, sim bool,
	) (newCtx sdk.Context, err error) {
		var anteHandler sdk.AnteHandler

		defer Recover(ctx.Logger(), &err)
		txWithExtensions, ok := tx.(authante.HasExtensionOptionsTx)
		if ok {
			opts := txWithExtensions.GetExtensionOptions()
			if len(opts) > 0 {
				switch typeURL := opts[0].GetTypeUrl(); typeURL {
				case "/ethermint.evm.v1.ExtensionOptionsEthereumTx":
					// handle as *evmtypes.MsgEthereumTx
					anteHandler = sdk.ChainAnteDecorators(
						evmante.NewEthSetUpContextDecorator(options.EvmKeeper), // outermost AnteDecorator. SetUpContext must be called first
						evmante.NewEthMempoolFeeDecorator(options.EvmKeeper),   // Check eth effective gas price against minimal-gas-prices
						evmante.NewEthValidateBasicDecorator(options.EvmKeeper),
						evmante.NewEthSigVerificationDecorator(options.EvmKeeper),
						evmante.NewEthAccountVerificationDecorator(options.AccountKeeper, options.BankKeeper, options.EvmKeeper),
						evmante.NewEthGasConsumeDecorator(options.EvmKeeper, options.MaxTxGasWanted),
						evmante.NewCanTransferDecorator(options.EvmKeeper),
						evmante.NewEthIncrementSenderSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator.
						permissionmodulekeeper.NewAuthDecorator(options.PermissionKeeper),
					)
				case "/ethermint.types.v1.ExtensionOptionsWeb3Tx":
					// handle as normal Cosmos SDK tx, except signature is checked for EIP712 representation
					anteHandler = sdk.ChainAnteDecorators(
						evmante.RejectMessagesDecorator{}, // reject MsgEthereumTxs
						authante.NewSetUpContextDecorator(),
						// NOTE: extensions option decorator removed
						// ante.NewRejectExtensionOptionsDecorator(),
						authante.NewMempoolFeeDecorator(),
						authante.NewValidateBasicDecorator(),
						authante.NewTxTimeoutHeightDecorator(),
						authante.NewValidateMemoDecorator(options.AccountKeeper),
						authante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
						authante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
						// SetPubKeyDecorator must be called before all signature verification decorators
						authante.NewSetPubKeyDecorator(options.AccountKeeper),
						authante.NewValidateSigCountDecorator(options.AccountKeeper),
						authante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
						// Note: signature verification uses EIP instead of the cosmos signature validator
						evmante.NewEip712SigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
						authante.NewIncrementSequenceDecorator(options.AccountKeeper),
						ibcante.NewAnteDecorator(options.IBCKeeper),
						wasmmodulekeeper.NewLimitSimulationGasDecorator(options.WasmConfig.SimulationGasLimit), // after setup context to enforce limits early
						wasmmodulekeeper.NewCountTXDecorator(options.TXCounterStoreKey),
						permissionmodulekeeper.NewAuthDecorator(options.PermissionKeeper),
					)
				default:
					return ctx, sdkerrors.Wrapf(
						sdkerrors.ErrUnknownExtensionOptions,
						"rejecting tx with unsupported extension option: %s", typeURL,
					)
				}

				return anteHandler(ctx, tx, sim)
			}
		}

		// handle as totally normal Cosmos SDK tx
		switch tx.(type) {
		case sdk.Tx:
			anteHandler = sdk.ChainAnteDecorators(
				evmante.RejectMessagesDecorator{}, // reject MsgEthereumTxs
				authante.NewSetUpContextDecorator(),
				authante.NewRejectExtensionOptionsDecorator(),
				authante.NewMempoolFeeDecorator(),
				authante.NewValidateBasicDecorator(),
				authante.NewTxTimeoutHeightDecorator(),
				authante.NewValidateMemoDecorator(options.AccountKeeper),
				authante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
				authante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
				// SetPubKeyDecorator must be called before all signature verification decorators
				authante.NewSetPubKeyDecorator(options.AccountKeeper),
				authante.NewValidateSigCountDecorator(options.AccountKeeper),
				authante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
				authante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
				authante.NewIncrementSequenceDecorator(options.AccountKeeper),
				ibcante.NewAnteDecorator(options.IBCKeeper),
				wasmmodulekeeper.NewLimitSimulationGasDecorator(options.WasmConfig.SimulationGasLimit), // after setup context to enforce limits early
				wasmmodulekeeper.NewCountTXDecorator(options.TXCounterStoreKey),
				permissionmodulekeeper.NewAuthDecorator(options.PermissionKeeper),
			)
		default:
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type: %T", tx)
		}

		return anteHandler(ctx, tx, sim)
	}, nil
}

func Recover(logger tmlog.Logger, err *error) {
	if r := recover(); r != nil {
		*err = sdkerrors.Wrapf(sdkerrors.ErrPanic, "%v", r)

		if e, ok := r.(error); ok {
			logger.Error(
				"ante handler panicked",
				"error", e,
				"stack trace", string(debug.Stack()),
			)
		} else {
			logger.Error(
				"ante handler panicked",
				"recover", fmt.Sprintf("%v", r),
			)
		}
	}
}
