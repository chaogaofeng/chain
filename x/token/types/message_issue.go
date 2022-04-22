package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"math/big"
)

const TypeMsgIssue = "issue"

var _ sdk.Msg = &MsgIssue{}

func NewMsgIssue(creator string, base string, display string, exp uint32, name, symbol, desc string, amount string, to string) *MsgIssue {
	return &MsgIssue{
		Creator:  creator,
		Base:     base,
		Display:  display,
		Exponent: exp,
		Symbol:   symbol,
		Name:     name,
		Desc:     desc,
		Amount:   amount,
		To:       to,
	}
}

func (msg *MsgIssue) Route() string {
	return RouterKey
}

func (msg *MsgIssue) Type() string {
	return TypeMsgIssue
}

func (msg *MsgIssue) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgIssue) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIssue) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	metaData := &banktypes.Metadata{
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
	if err := metaData.Validate(); err != nil {
		return err
	}

	if amt, ok := new(big.Int).SetString(msg.Amount, 10); !ok || amt.Cmp(big.NewInt(0)) <= 0 {
		return sdkerrors.Wrapf(ErrInvalidAmount, "invalid token amount %s", msg.Amount)
	}

	if len(msg.To) > 0 {
		_, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid to address (%s)", err)
		}
	}
	return nil
}
