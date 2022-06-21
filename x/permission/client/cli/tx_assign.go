package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/goldnet/chain/x/permission/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAssign() *cobra.Command {
	cmd := &cobra.Command{
		Use: "assign [address] [roles]",
		Short: fmt.Sprintf("Assign roles to an address. Auth options: %s", func(m map[string]int32) []string {
			j := 0
			keys := make([]string, len(m))
			for k := range m {
				keys[j] = k
				j++
			}
			return keys
		}(types.Role_value)),
		Args: cobra.MinimumNArgs(2),
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

			msg := types.NewMsgAssign(
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
