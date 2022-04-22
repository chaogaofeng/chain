package cli

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/glodnet/chain/x/identity/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing identity",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			certFile := viper.GetString(FlagCertificateFile)
			var cert []byte
			if len(certFile) > 0 {
				cert, err = ioutil.ReadFile(certFile)
				if err != nil {
					return fmt.Errorf("failed to read the certificate file: %s", err.Error())
				}
			}

			msg := types.NewMsgUpdate(
				clientCtx.GetFromAddress().String(),
				string(cert),
				viper.GetString(FlagParent),
				viper.GetString(FlagData),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagCertificateFile, "", "X.509 certificate file path")
	cmd.Flags().String(FlagParent, "", "address of parent certificate")
	cmd.Flags().String(FlagData, "", "custom data of the identity")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
