package keeper_test

import (
	"testing"

	testkeeper "github.com/glodnet/chain/testutil/keeper"
	"github.com/glodnet/chain/x/token/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TokenKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.Enabled, k.Enabled(ctx))
}
