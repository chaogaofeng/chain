package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/glodnet/chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgRemove_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemove
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRemove{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRemove{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
