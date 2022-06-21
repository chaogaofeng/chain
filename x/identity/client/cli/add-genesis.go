package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/goldnet/chain/pkg/cacmd/ca"
	"github.com/goldnet/chain/x/identity/types"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
)

// AddGenesisCmd returns add-genesis-account-id cobra Command.
func AddGenesisCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-identity",
		Short: "Add a genesis account identity to genesis.json",
		Long: `Add a genesis account identity to genesis.json. The provided account must specify
the account address or key name and an identity. If a key name is given,
the address will be looked up in the local Keybase. 
`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			certFile := viper.GetString(FlagCertificateFile)
			var certBytes []byte
			if len(certFile) > 0 {
				bts, err := ioutil.ReadFile(certFile)
				if err != nil {
					return fmt.Errorf("failed to read the certificate file: %s", err.Error())
				}
				certBytes = bts
			}

			msg := types.NewMsgCreate(
				clientCtx.GetFromAddress().String(),
				string(certBytes),
				viper.GetString(FlagParent),
				viper.GetString(FlagData),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			cert, _ := ca.ReadCertificateFromMem([]byte(msg.Certificate))
			pubKey, _ := cert.GetPubkeyFromCert()
			addr := sdk.AccAddress(pubKey.Address())
			identity := types.Identity{
				Address:     addr.String(),
				Certificate: msg.Certificate,
				Parent:      msg.Parent,
				Data:        msg.Data,
			}

			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			// identity genesis
			var genesisState types.GenesisState
			if appState[types.ModuleName] != nil {
				cdc.MustUnmarshalJSON(appState[types.ModuleName], &genesisState)
			}
			genesisState.IdentityList = append(genesisState.IdentityList, identity)

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

			// completed
			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}
			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(FlagCertificateFile, "", "X.509 certificate file path")
	cmd.Flags().String(FlagParent, "", "address of parent certificate")
	cmd.Flags().String(FlagData, "", "custom data of the identity")
	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test)")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
