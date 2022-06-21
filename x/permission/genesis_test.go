package permission_test

import (
	"testing"

	keepertest "github.com/goldnet/chain/testutil/keeper"
	"github.com/goldnet/chain/testutil/nullify"
	"github.com/goldnet/chain/x/permission"
	"github.com/goldnet/chain/x/permission/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RoleAccountList: []types.RoleAccount{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PermissionKeeper(t)
	permission.InitGenesis(ctx, *k, genesisState)
	got := permission.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RoleAccountList, got.RoleAccountList)
	// this line is used by starport scaffolding # genesis/test/assert
}
