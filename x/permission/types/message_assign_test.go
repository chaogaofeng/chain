package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/goldnet/chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAssign_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAssign
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAssign{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAssign{
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
