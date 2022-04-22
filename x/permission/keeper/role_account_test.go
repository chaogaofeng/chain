package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/glodnet/chain/testutil/keeper"
	"github.com/glodnet/chain/testutil/nullify"
	"github.com/glodnet/chain/x/permission/keeper"
	"github.com/glodnet/chain/x/permission/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRoleAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RoleAccount {
	items := make([]types.RoleAccount, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetRoleAccount(ctx, items[i])
	}
	return items
}

func TestRoleAccountGet(t *testing.T) {
	keeper, ctx := keepertest.PermissionKeeper(t)
	items := createNRoleAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRoleAccount(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRoleAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.PermissionKeeper(t)
	items := createNRoleAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRoleAccount(ctx,
			item.Address,
		)
		_, found := keeper.GetRoleAccount(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestRoleAccountGetAll(t *testing.T) {
	keeper, ctx := keepertest.PermissionKeeper(t)
	items := createNRoleAccount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRoleAccount(ctx)),
	)
}
