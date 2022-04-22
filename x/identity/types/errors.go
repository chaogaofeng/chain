package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/identity module sentinel errors
var (
	ErrInvalidCertificate         = sdkerrors.Register(ModuleName, 2, "invalid certificate")
	ErrIdentityExists             = sdkerrors.Register(ModuleName, 3, "identity already exists")
	ErrUnknownIdentity            = sdkerrors.Register(ModuleName, 4, "unknown identity")
	ErrUnsupportedPubKeyAlgorithm = sdkerrors.Register(ModuleName, 5, "unsupported public key algorithm; only RSA, DSA, ECDSA, ED25519 and SM2 supported")
	ErrNotAuthorized              = sdkerrors.Register(ModuleName, 6, "owner not matching")
	ErrInvalidNodeID              = sdkerrors.Register(ModuleName, 7, "invalid node ID")
)
