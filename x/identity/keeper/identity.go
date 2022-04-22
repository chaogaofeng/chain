package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glodnet/chain/x/identity/types"
)

// SetIdentity set a specific identity in the store from its index
func (k Keeper) SetIdentity(ctx sdk.Context, identity types.Identity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKeyPrefix))
	b := k.cdc.MustMarshal(&identity)
	store.Set(types.IdentityKey(
		identity.Address,
	), b)
}

// GetIdentity returns a identity from its index
func (k Keeper) GetIdentity(
	ctx sdk.Context,
	address string,

) (val types.Identity, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKeyPrefix))

	b := store.Get(types.IdentityKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveIdentity removes a identity from the store
func (k Keeper) RemoveIdentity(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKeyPrefix))
	store.Delete(types.IdentityKey(
		address,
	))
}

// GetAllIdentity returns all identity
func (k Keeper) GetAllIdentity(ctx sdk.Context) (list []types.Identity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Identity
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
