package identity_test

import (
	"testing"

	keepertest "github.com/goldnet/chain/testutil/keeper"
	"github.com/goldnet/chain/testutil/nullify"
	"github.com/goldnet/chain/x/identity"
	"github.com/goldnet/chain/x/identity/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		IdentityList: []types.Identity{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityKeeper(t)
	identity.InitGenesis(ctx, *k, genesisState)
	got := identity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.IdentityList, got.IdentityList)
	// this line is used by starport scaffolding # genesis/test/assert
}
