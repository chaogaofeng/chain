package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/goldnet/chain/x/permission/types"
)

// SetRoleAccount set a specific roleAccount in the store from its index
func (k Keeper) SetRoleAccount(ctx sdk.Context, roleAccount types.RoleAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RoleAccountKeyPrefix))
	b := k.cdc.MustMarshal(&roleAccount)
	store.Set(types.RoleAccountKey(
		roleAccount.Address,
	), b)
}

// GetRoleAccount returns a roleAccount from its index
func (k Keeper) GetRoleAccount(
	ctx sdk.Context,
	address string,

) (val types.RoleAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RoleAccountKeyPrefix))

	b := store.Get(types.RoleAccountKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRoleAccount removes a roleAccount from the store
func (k Keeper) RemoveRoleAccount(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RoleAccountKeyPrefix))
	store.Delete(types.RoleAccountKey(
		address,
	))
}

// GetAllRoleAccount returns all roleAccount
func (k Keeper) GetAllRoleAccount(ctx sdk.Context) (list []types.RoleAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RoleAccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RoleAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
