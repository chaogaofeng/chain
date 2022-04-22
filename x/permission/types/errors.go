package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/permission module sentinel errors
var (
	ErrUnauthorizedOperation = sdkerrors.Register(ModuleName, 2, "unauthorized operation")
	ErrRemoveUnknownRole     = sdkerrors.Register(ModuleName, 3, "the account does not have this role")
	ErrInvalidMsgURL         = sdkerrors.Register(ModuleName, 4, "invalid url")
)
