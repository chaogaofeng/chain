package types

import (
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/goldnet/chain/pkg/cacmd/ca"
)

const TypeMsgCreate = "create"

var _ sdk.Msg = &MsgCreate{}

func NewMsgCreate(creator string, certificate string, parent string, data string) *MsgCreate {
	return &MsgCreate{
		Creator:     creator,
		Certificate: certificate,
		Parent:      parent,
		Data:        data,
	}
}

func (msg *MsgCreate) Route() string {
	return RouterKey
}

func (msg *MsgCreate) Type() string {
	return TypeMsgCreate
}

func (msg *MsgCreate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	//if len(msg.Certificate) > 0 {
	//if err := CheckCertificate([]byte(msg.Certificate)); err != nil {
	//	return err
	//}

	cert, err := ca.ReadCertificateFromMem([]byte(msg.Certificate))
	if err != nil {
		return err
	}
	pk, err := cert.GetPubkeyFromCert()
	if err != nil {
		return err
	}
	if _, err := cryptocodec.FromTmPubKeyInterface(pk); err != nil {
		return err
	}
	//}

	if len(msg.Parent) > 0 {
		_, err := sdk.AccAddressFromBech32(msg.Parent)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid parent address (%s)", err)
		}
	}
	return nil
}
