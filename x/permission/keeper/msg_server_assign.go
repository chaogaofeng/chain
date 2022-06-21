package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/telemetry"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/goldnet/chain/x/permission/types"
)

func (k msgServer) Assign(goCtx context.Context, msg *types.MsgAssign) (*types.MsgAssignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, _ := sdk.AccAddressFromBech32(msg.Address)
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	if err := k.Authorize(ctx, addr.String(), creator.String(), msg.Roles...); err != nil {
		return nil, err
	}

	// Create account if addr does not exist.
	acc := k.accountKeeper.GetAccount(ctx, addr)
	if acc == nil {
		defer telemetry.IncrCounter(1, "new", "account")
		k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccountWithAddress(ctx, addr))
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
	return &types.MsgAssignResponse{}, nil
}
