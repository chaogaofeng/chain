package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnassign = "unassign"

var _ sdk.Msg = &MsgUnassign{}

func NewMsgUnassign(creator string, address string, roles []Role) *MsgUnassign {
	return &MsgUnassign{
		Creator: creator,
		Address: address,
		Roles:   roles,
	}
}

func (msg *MsgUnassign) Route() string {
	return RouterKey
}

func (msg *MsgUnassign) Type() string {
	return TypeMsgUnassign
}

func (msg *MsgUnassign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnassign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnassign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid address")
	}

	if len(msg.Roles) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "roles missing")
	}
	for _, r := range msg.Roles {
		if !ValidRole(r) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid role %s", r.String())
		}
	}

	return nil
}
