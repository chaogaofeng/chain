package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/token module sentinel errors
var (
	ErrInvalidBase   = sdkerrors.Register(ModuleName, 2, "invalid base denom")
	ErrInvalidAmount = sdkerrors.Register(ModuleName, 6, "invalid amount")

	ErrTokenAlreadyExist = sdkerrors.Register(ModuleName, 10, "token alreay exists")
	ErrTokenNotExist     = sdkerrors.Register(ModuleName, 11, "token does not exist")
	ErrInvalidOwner      = sdkerrors.Register(ModuleName, 12, "invalid token owner")
)
