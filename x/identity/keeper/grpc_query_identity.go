package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/goldnet/chain/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IdentityAll(c context.Context, req *types.QueryAllIdentityRequest) (*types.QueryAllIdentityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var identitys []types.Identity
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	identityStore := prefix.NewStore(store, types.KeyPrefix(types.IdentityKeyPrefix))

	pageRes, err := query.Paginate(identityStore, req.Pagination, func(key []byte, value []byte) error {
		var identity types.Identity
		if err := k.cdc.Unmarshal(value, &identity); err != nil {
			return err
		}

		identitys = append(identitys, identity)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIdentityResponse{Identity: identitys, Pagination: pageRes}, nil
}

func (k Keeper) Identity(c context.Context, req *types.QueryGetIdentityRequest) (*types.QueryGetIdentityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetIdentity(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetIdentityResponse{Identity: val}, nil
}
