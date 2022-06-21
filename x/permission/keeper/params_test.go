package keeper_test

import (
	"testing"

	testkeeper "github.com/goldnet/chain/testutil/keeper"
	"github.com/goldnet/chain/x/permission/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PermissionKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.Enabled, k.Enabled(ctx))
}
