package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math/big"
)

const TypeMsgBurn = "burn"

var _ sdk.Msg = &MsgBurn{}

func NewMsgBurn(creator string, base string, amount string) *MsgBurn {
	return &MsgBurn{
		Creator: creator,
		Base:    base,
		Amount:  amount,
	}
}

func (msg *MsgBurn) Route() string {
	return RouterKey
}

func (msg *MsgBurn) Type() string {
	return TypeMsgBurn
}

func (msg *MsgBurn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := sdk.ValidateDenom(msg.Base); err != nil {
		return sdkerrors.Wrapf(ErrInvalidBase, "invalid token base denom: %w", err)
	}

	if amt, ok := new(big.Int).SetString(msg.Amount, 10); !ok || amt.Cmp(big.NewInt(0)) <= 0 {
		return sdkerrors.Wrapf(ErrInvalidAmount, "invalid token amount %s", msg.Amount)
	}

	return nil
}
