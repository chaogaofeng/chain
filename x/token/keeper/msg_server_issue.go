package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/goldnet/chain/x/token/types"
)

func (k msgServer) Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.bankKeeper.GetDenomMetaData(ctx, msg.Base); found {
		return nil, sdkerrors.Wrapf(types.ErrTokenAlreadyExist, msg.Base)
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

	denomMetaData := banktypes.Metadata{
		Description: msg.Desc,
		Base:        msg.Base,
		Display:     msg.Display,
		Symbol:      msg.Symbol,
		Name:        msg.Name,
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    msg.Base,
				Exponent: 0,
			},
			{
				Denom:    msg.Display,
				Exponent: msg.Exponent,
			},
		},
	}
	token := types.Token{
		Symbol:   msg.Base,
		Owner:    msg.Creator,
		Metadata: &denomMetaData,
		Issued:   amountCoin,
		Burned:   sdk.NewCoin(msg.Base, sdk.ZeroInt()),
	}
	k.SetToken(ctx, token)
	k.bankKeeper.SetDenomMetaData(ctx, denomMetaData)

	// mint coins into module account
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amountCoin)); err != nil {
		return nil, err
	}

	// sent coins to owner's account
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, sdk.NewCoins(amountCoin)); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			strings.Join([]string{msg.Route(), msg.Type()}, "_"),
			sdk.NewAttribute(types.AttributeKeyBase, msg.Base),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amountCoin.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgIssueResponse{}, nil
}
