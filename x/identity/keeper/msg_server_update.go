package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/goldnet/chain/pkg/cacmd/ca"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/goldnet/chain/x/identity/types"
)

func (k msgServer) Update(goCtx context.Context, msg *types.MsgUpdate) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cert, _ := ca.ReadCertificateFromMem([]byte(msg.Certificate))
	pubKey, _ := cert.GetPubkeyFromCert()
	address := sdk.AccAddress(pubKey.Address()).String()

	if identity, found := k.GetIdentity(ctx, address); !found {
		return nil, sdkerrors.Wrapf(types.ErrUnknownIdentity, address)
	} else if identity.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAuthorized, identity.Creator)
	}
	if _, err := k.VerifyCertificate(ctx, address, msg.Certificate, msg.Parent); err != nil {
		return nil, err
	}

	k.SetIdentity(ctx, types.Identity{
		Address:     address,
		Certificate: msg.Certificate,
		Parent:      msg.Parent,
		Data:        msg.Data,
		Creator:     msg.Creator,
	})

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			msg.Type(),
			sdk.NewAttribute(types.AttributeKeyAddress, address),
			sdk.NewAttribute(types.AttributeKeyParent, msg.Parent),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgUpdateResponse{
		Address: address,
	}, nil
}
