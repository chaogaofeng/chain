package cli

import (
	"strconv"

	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/goldnet/chain/x/token/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [base denom]",
		Short: "Mint a existing token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			base := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(
				clientCtx.GetFromAddress().String(),
				base,
				viper.GetString(FlagAmount),
				viper.GetString(FlagTo),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagAmount, "0", "Amount of the token to be minted")
	cmd.Flags().String(FlagTo, "", "Address to which the token is to be minted")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
