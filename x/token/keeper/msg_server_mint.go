package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glodnet/chain/x/token/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token, found := k.GetToken(ctx, msg.Base)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTokenNotExist, msg.Base)
	} else if token.Owner != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the token %s", msg.Creator, msg.Base)
	}

	owner, _ := sdk.AccAddressFromBech32(msg.Creator)
	var recipient sdk.AccAddress
	if len(msg.To) != 0 {
		recipient, _ = sdk.AccAddressFromBech32(msg.To)
	} else {
		recipient = owner
	}

	amt, _ := sdk.NewIntFromString(msg.Amount)
	amountCoin := sdk.NewCoin(msg.Base, amt)

	// mint coins into module account
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amountCoin)); err != nil {
		return nil, err
	}

	// sent coins to owner's account
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, sdk.NewCoins(amountCoin)); err != nil {
		return nil, err
	}

	token.Issued = token.Issued.Add(amountCoin)
	k.SetToken(ctx, token)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			strings.Join([]string{msg.Route(), msg.Type()}, "_"),
			sdk.NewAttribute(types.AttributeKeyBase, msg.Base),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amountCoin.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgMintResponse{}, nil
}
