package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/goldnet/chain/pkg/cacmd/ca"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/goldnet/chain/x/identity/types"
)

func (k msgServer) Create(goCtx context.Context, msg *types.MsgCreate) (*types.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cert, _ := ca.ReadCertificateFromMem([]byte(msg.Certificate))
	pubKey, _ := cert.GetPubkeyFromCert()
	addr := sdk.AccAddress(pubKey.Address())
	address := addr.String()

	if _, found := k.GetIdentity(ctx, address); found {
		return nil, sdkerrors.Wrapf(types.ErrIdentityExists, address)
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

	// Create account if addr does not exist.
	acc := k.accountKeeper.GetAccount(ctx, addr)
	if acc == nil {
		defer telemetry.IncrCounter(1, "new", "account")
		k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccountWithAddress(ctx, addr))
	}

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

	return &types.MsgCreateResponse{
		Address: address,
	}, nil
}
