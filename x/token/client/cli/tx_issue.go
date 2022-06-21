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

func CmdIssue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue [base denom]",
		Short: "Issue a new token. The base denom, minimum unit name of the token.(eg: uatom)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			base := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssue(
				clientCtx.GetFromAddress().String(),
				base,
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
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagName, "", "name defines the name of the token (eg: Cosmos Atom)")
	cmd.Flags().String(FlagSymbol, "", "symbol is the token symbol usually shown on exchanges (eg: ATOM).")
	cmd.Flags().String(FlagDisplay, "", "display is the token suggested denom that should be displayed in clients (eg: atom).")
	cmd.Flags().Uint32(FlagExponent, 0, "exponent represents power of 10 exponent that display must raise the base (eg: 1 atom = 10^6 uatom).")
	cmd.Flags().String(FlagDescription, "", "The token description.")
	cmd.Flags().String(FlagAmount, "0", "Amount of the token to be issued")
	cmd.Flags().String(FlagTo, "", "Address to which the token is to be issued")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
