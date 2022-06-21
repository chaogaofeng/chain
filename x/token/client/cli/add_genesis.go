package cli

import (
	"bufio"
	"encoding/json"
	"fmt"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/goldnet/chain/x/token/types"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
)

// AddGenesisCmd returns add-genesis cobra Command.
func AddGenesisCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-token [address_or_key_name] [base denom]",
		Short: "Add a genesis token to genesis.json",
		Long: `Add a genesis token to genesis.json. The provided account must specify
the account address or key name and a token. If a key name is given,
the address will be looked up in the local Keybase.
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			var kr keyring.Keyring
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				inBuf := bufio.NewReader(cmd.InOrStdin())
				keyringBackend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)
				if keyringBackend != "" && clientCtx.Keyring == nil {
					var err error
					kr, err = keyring.New(sdk.KeyringServiceName(), keyringBackend, clientCtx.HomeDir, inBuf)
					if err != nil {
						return err
					}
				} else {
					kr = clientCtx.Keyring
				}

				info, err := kr.Key(args[0])
				if err != nil {
					return fmt.Errorf("failed to get address from Keyring: %w", err)
				}
				addr = info.GetAddress()
			}

			msg := types.NewMsgIssue(
				addr.String(),
				args[1],
				viper.GetString(FlagDisplay),
				viper.GetUint32(FlagExponent),
				viper.GetString(FlagName),
				viper.GetString(FlagSymbol),
				viper.GetString(FlagDescription),
				viper.GetString(FlagAmount),
				viper.GetString(FlagTo),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			var recipient sdk.AccAddress
			if len(msg.To) != 0 {
				recipient, _ = sdk.AccAddressFromBech32(msg.To)
			} else {
				recipient = addr
			}
			amt, _ := sdk.NewIntFromString(msg.Amount)
			amountCoin := sdk.NewCoin(msg.Base, amt)

			denomMetaData := banktypes.Metadata{
				Description: msg.Desc,
				Base:        msg.Base,
				Display:     msg.Display,
				Symbol:      msg.Symbol,
				Name:        msg.Name,
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    msg.Base,
						Exponent: 0,
					},
					{
						Denom:    msg.Display,
						Exponent: msg.Exponent,
					},
				},
			}
			token := types.Token{
				Symbol:   msg.Base,
				Owner:    msg.Creator,
				Metadata: &denomMetaData,
				Issued:   amountCoin,
				Burned:   sdk.NewCoin(msg.Base, sdk.ZeroInt()),
			}

			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			// token genesis
			var genesisState types.GenesisState
			if appState[types.ModuleName] != nil {
				cdc.MustUnmarshalJSON(appState[types.ModuleName], &genesisState)
			}
			genesisState.TokenList = append(genesisState.TokenList, token)
			genesisStateBz, err := cdc.MarshalJSON(&genesisState)
			if err != nil {
				return fmt.Errorf("failed to marshal genesis state: %w", err)
			}
			appState[types.ModuleName] = genesisStateBz

			// auth genesis
			authGenState := authtypes.GetGenesisStateFromAppState(cdc, appState)
			accs, err := authtypes.UnpackAccounts(authGenState.Accounts)
			if err != nil {
				return fmt.Errorf("failed to get accounts from any: %w", err)
			}
			if !accs.Contains(addr) {
				genAccount := authtypes.NewBaseAccount(addr, nil, 0, 0)
				accs = append(accs, genAccount)
			}
			if !accs.Contains(recipient) {
				genAccount := authtypes.NewBaseAccount(recipient, nil, 0, 0)
				accs = append(accs, genAccount)
			}
			accs = authtypes.SanitizeGenesisAccounts(accs)
			genAccs, err := authtypes.PackAccounts(accs)
			if err != nil {
				return fmt.Errorf("failed to convert accounts into any's: %w", err)
			}
			authGenState.Accounts = genAccs
			authGenStateBz, err := cdc.MarshalJSON(&authGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal auth genesis state: %w", err)
			}
			appState[authtypes.ModuleName] = authGenStateBz

			// bank genesis
			bankGenState := banktypes.GetGenesisStateFromAppState(cdc, appState)
			bankGenState.Balances = append(bankGenState.Balances, banktypes.Balance{Address: recipient.String(), Coins: sdk.NewCoins(amountCoin).Sort()})
			bankGenState.Balances = banktypes.SanitizeGenesisBalances(bankGenState.Balances)
			bankGenState.DenomMetadata = append(bankGenState.DenomMetadata, denomMetaData)
			bankGenStateBz, err := cdc.MarshalJSON(bankGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal bank genesis state: %w", err)
			}
			appState[banktypes.ModuleName] = bankGenStateBz

			// completed
			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}
			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test)")
	cmd.Flags().String(FlagName, "", "name defines the name of the token (eg: Cosmos Atom)")
	cmd.Flags().String(FlagSymbol, "", "symbol is the token symbol usually shown on exchanges (eg: ATOM).")
	cmd.Flags().String(FlagDisplay, "", "display is the token suggested denom that should be displayed in clients (eg: atom).")
	cmd.Flags().Uint32(FlagExponent, 0, "exponent represents power of 10 exponent that display must raise the base (eg: 1 atom = 10^6 uatom).")
	cmd.Flags().String(FlagDescription, "", "The token description.")
	cmd.Flags().String(FlagAmount, "0", "Amount of the token to be issued")
	cmd.Flags().String(FlagTo, "", "Address to which the token is to be issued")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
