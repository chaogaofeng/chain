package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glodnet/chain/x/token/types"
)

// SetToken set a specific token in the store from its index
func (k Keeper) SetToken(ctx sdk.Context, token types.Token) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKeyPrefix))
	b := k.cdc.MustMarshal(&token)
	store.Set(types.TokenKey(
		token.Symbol,
	), b)
}

// GetToken returns a token from its index
func (k Keeper) GetToken(
	ctx sdk.Context,
	symbol string,

) (val types.Token, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKeyPrefix))

	b := store.Get(types.TokenKey(
		symbol,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveToken removes a token from the store
func (k Keeper) RemoveToken(
	ctx sdk.Context,
	symbol string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKeyPrefix))
	store.Delete(types.TokenKey(
		symbol,
	))
}

// GetAllToken returns all token
func (k Keeper) GetAllToken(ctx sdk.Context) (list []types.Token) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
