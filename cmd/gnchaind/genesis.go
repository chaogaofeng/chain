package main

import (
	identitycli "github.com/goldnet/chain/x/identity/client/cli"
	permissioncli "github.com/goldnet/chain/x/permission/client/cli"
	tokencli "github.com/goldnet/chain/x/token/client/cli"
	"github.com/spf13/cobra"
)

func genesisCommand(defaultNodeHome string) *cobra.Command {
	c := &cobra.Command{
		Use:   "genesis",
		Short: "add init data to genesis.json",
	}
	c.AddCommand(permissioncli.AddGenesisCmd(defaultNodeHome))
	c.AddCommand(identitycli.AddGenesisCmd(defaultNodeHome))
	c.AddCommand(tokencli.AddGenesisCmd(defaultNodeHome))
	return c
}
