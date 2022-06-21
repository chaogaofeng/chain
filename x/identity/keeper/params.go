package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/goldnet/chain/x/identity/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Enabled(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// Enabled returns the Enabled param
func (k Keeper) Enabled(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyEnabled, &res)
	return
}
