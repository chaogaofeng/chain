package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glodnet/chain/x/identity/types"
)

func (k msgServer) Remove(goCtx context.Context, msg *types.MsgRemove) (*types.MsgRemoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if identity, found := k.GetIdentity(ctx, msg.Address); !found {
		return nil, sdkerrors.Wrapf(types.ErrUnknownIdentity, msg.Address)
	} else if identity.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAuthorized, identity.Creator)
	}

	k.RemoveIdentity(ctx, msg.Address)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			msg.Type(),
			sdk.NewAttribute(types.AttributeKeyAddress, msg.Address),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgRemoveResponse{}, nil
}
