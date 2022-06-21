package token_test

import (
	"testing"

	keepertest "github.com/goldnet/chain/testutil/keeper"
	"github.com/goldnet/chain/testutil/nullify"
	"github.com/goldnet/chain/x/token"
	"github.com/goldnet/chain/x/token/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TokenList: []types.Token{
			{
				Symbol: "0",
			},
			{
				Symbol: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenKeeper(t)
	token.InitGenesis(ctx, *k, genesisState)
	got := token.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TokenList, got.TokenList)
	// this line is used by starport scaffolding # genesis/test/assert
}
