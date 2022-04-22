package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/glodnet/chain/x/permission/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RoleAccountAll(c context.Context, req *types.QueryAllRoleAccountRequest) (*types.QueryAllRoleAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var roleAccounts []types.RoleAccount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	roleAccountStore := prefix.NewStore(store, types.KeyPrefix(types.RoleAccountKeyPrefix))

	pageRes, err := query.Paginate(roleAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var roleAccount types.RoleAccount
		if err := k.cdc.Unmarshal(value, &roleAccount); err != nil {
			return err
		}

		roleAccounts = append(roleAccounts, roleAccount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRoleAccountResponse{RoleAccount: roleAccounts, Pagination: pageRes}, nil
}

func (k Keeper) RoleAccount(c context.Context, req *types.QueryGetRoleAccountRequest) (*types.QueryGetRoleAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRoleAccount(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetRoleAccountResponse{RoleAccount: val}, nil
}
