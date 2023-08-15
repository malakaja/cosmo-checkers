package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/malakaja/cosmo-checkers/checkers/testutil/sample"
	"github.com/malakaja/cosmo-checkers/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateGame_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateGame
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgCreateGame{
				Creator: "invalid_address",
				Black:   sample.AccAddress(),
				Red:     sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgCreateGame{
				Creator: sample.AccAddress(),
				Black:   sample.AccAddress(),
				Red:     sample.AccAddress(),
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
