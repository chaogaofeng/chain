package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/goldnet/chain/x/token/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token, found := k.GetToken(ctx, msg.Base)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTokenNotExist, msg.Base)
	}

	owner, _ := sdk.AccAddressFromBech32(msg.Creator)
	amt, _ := sdk.NewIntFromString(msg.Amount)
	amountCoin := sdk.NewCoin(msg.Base, amt)

	// burn coins
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, sdk.NewCoins(amountCoin)); err != nil {
		return nil, err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amountCoin)); err != nil {
		return nil, err
	}

	token.Burned = token.Burned.Add(amountCoin)
	k.SetToken(ctx, token)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			strings.Join([]string{msg.Route(), msg.Type()}, "_"),
			sdk.NewAttribute(types.AttributeKeyBase, msg.Base),
			sdk.NewAttribute(types.AttributeKeyAmount, amountCoin.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgBurnResponse{}, nil
}
