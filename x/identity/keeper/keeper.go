package keeper

import (
	"encoding/hex"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/goldnet/chain/pkg/cacmd/ca"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/goldnet/chain/x/identity/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) VerifyCertificate(ctx sdk.Context, address string, certificate string, parent string) (ca.Cert, error) {
	cert, _ := ca.ReadCertificateFromMem([]byte(certificate))
	if len(parent) == 0 {
		identity, found := k.GetIdentity(ctx, parent)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrUnknownIdentity, parent)
		}

		parentCert, err := k.VerifyCertificate(ctx, identity.Address, identity.Certificate, identity.Parent)
		if err != nil {
			return nil, err
		}

		if err = cert.VerifyCertFromRoot(parentCert); err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidCertificate, "cannot be verified by parent certificate, err: %s", err)
		}
	}
	return cert, nil
}

// FilterNodeByID implements sdk.PeerFilter
func (k Keeper) FilterNodeByID(ctx sdk.Context, nodeID string) abci.ResponseQuery {
	if !k.GetParams(ctx).Enabled {
		return abci.ResponseQuery{}
	}

	id, err := hex.DecodeString(nodeID)
	if err != nil {
		return abci.ResponseQuery{
			Code: types.ErrInvalidNodeID.ABCICode(),
		}
	}

	address := sdk.AccAddress(id).String()
	identity, found := k.GetIdentity(ctx, address)
	if !found {
		return abci.ResponseQuery{
			Code: types.ErrUnknownIdentity.ABCICode(),
		}
	}
	if _, err := k.VerifyCertificate(ctx, address, identity.Certificate, identity.Parent); err != nil {
		return abci.ResponseQuery{
			Code: err.(*sdkerrors.Error).ABCICode(),
		}
	}

	return abci.ResponseQuery{}
}
