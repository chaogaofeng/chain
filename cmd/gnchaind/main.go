package main

import (
	"os"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	starportcmd "github.com/tendermint/starport/starport/cmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/glodnet/chain/app"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name+"_45-1",
		app.ModuleBasics,
		app.New,
		cosmoscmd.AddSubCmd(
			genesisCommand(app.DefaultNodeHome),
			testnetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}),
			chainCommand(),
			faucetServeCmd(),
			// caCommand(),
			starportcmd.NewRelayer(),
		),
		cosmoscmd.CustomizeStartCmd(app.CustomizeStartCmd),
	// this line is used by starport scaffolding # root/arguments
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
