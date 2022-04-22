package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/glodnet/chain/x/permission/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUnassign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unassign [address] [roles]",
		Short: "Unassign roles from an address",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			roles, err := types.GetRolesFromStr(args[1:]...)
			if err != nil {
				return err
			}

			msg := types.NewMsgUnassign(
				clientCtx.GetFromAddress().String(),
				argAddress,
				roles,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
