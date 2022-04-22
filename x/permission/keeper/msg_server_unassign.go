package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glodnet/chain/x/permission/types"
)

func (k msgServer) Unassign(goCtx context.Context, msg *types.MsgUnassign) (*types.MsgUnassignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, _ := sdk.AccAddressFromBech32(msg.Address)
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	if err := k.Unauthorize(ctx, addr.String(), creator.String(), msg.Roles...); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			strings.Join([]string{msg.Route(), msg.Type()}, "_"),
			sdk.NewAttribute(types.AttributeKeyAddress, msg.Address),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})
	return &types.MsgUnassignResponse{}, nil
}
