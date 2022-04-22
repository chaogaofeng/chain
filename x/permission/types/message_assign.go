package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAssign = "assign"

var _ sdk.Msg = &MsgAssign{}

func NewMsgAssign(creator string, address string, roles []Role) *MsgAssign {
	return &MsgAssign{
		Creator: creator,
		Address: address,
		Roles:   roles,
	}
}

func (msg *MsgAssign) Route() string {
	return RouterKey
}

func (msg *MsgAssign) Type() string {
	return TypeMsgAssign
}

func (msg *MsgAssign) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAssign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAssign) ValidateBasic() error {
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
