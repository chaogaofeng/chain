package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/glodnet/chain/x/permission/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
		AuthMap    map[string]types.Auth

		accountKeeper types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		AuthMap:       make(map[string]types.Auth),
		accountKeeper: accountKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Authorize assigns the specified roles to an address
func (k *Keeper) Authorize(ctx sdk.Context, address string, creator string, rs ...types.Role) error {
	addressAuth := types.AuthDefault
	if roleAccount, found := k.GetRoleAccount(ctx, address); found {
		addressAuth = types.FromRoles(roleAccount.Roles)
		if addressAuth.IsRootAdmin() {
			return sdkerrors.Wrap(types.ErrUnauthorizedOperation, "can not operate root admin")
		}
	}

	creatorAuth := types.AuthDefault
	if roleAccount, found := k.GetRoleAccount(ctx, creator); found {
		creatorAuth = types.FromRoles(roleAccount.Roles)
	}

	for _, r := range rs {
		if !creatorAuth.IsRootAdmin() && !creatorAuth.Access(r.Auth()) {
			return sdkerrors.Wrapf(types.ErrUnauthorizedOperation, "can not add %s role", r.String())
		}
		addressAuth = addressAuth | r.Auth()
	}

	k.SetRoleAccount(ctx, types.RoleAccount{
		Address: address,
		Roles:   addressAuth.Roles(),
	})
	return nil
}

// Unauthorize unassigns the specified roles from an address
func (k Keeper) Unauthorize(ctx sdk.Context, address string, creator string, roles ...types.Role) error {
	addressAuth := types.AuthDefault
	if roleAccount, found := k.GetRoleAccount(ctx, address); found {
		addressAuth = types.FromRoles(roleAccount.Roles)
	}

	creatorAuth := types.AuthDefault
	if roleAccount, found := k.GetRoleAccount(ctx, creator); found {
		creatorAuth = types.FromRoles(roleAccount.Roles)
	}

	for _, r := range roles {
		if !creatorAuth.IsRootAdmin() && !creatorAuth.Access(r.Auth()) {
			return sdkerrors.Wrapf(types.ErrUnauthorizedOperation, "can not remove %s role", r.String())
		}
		if !addressAuth.Access(r.Auth()) {
			return sdkerrors.Wrapf(types.ErrRemoveUnknownRole, "can not remove %s role", r.String())
		}
		addressAuth = addressAuth & (addressAuth ^ r.Auth())
	}
	if addressAuth == types.AuthDefault {
		k.RemoveRoleAccount(ctx, address)
	} else {
		k.SetRoleAccount(ctx, types.RoleAccount{
			Address: address,
			Roles:   addressAuth.Roles(),
		})
	}
	return nil
}

// Access checks the signer auth
func (k Keeper) Access(ctx sdk.Context, address string, auth types.Auth) error {
	roleAccount, _ := k.GetRoleAccount(ctx, address)
	accountAuth := types.FromRoles(roleAccount.Roles)
	if !auth.Access(accountAuth) {
		return sdkerrors.Wrapf(
			types.ErrUnauthorizedOperation,
			"Required roles: %s; sender roles: %s. ",
			auth.Roles(), accountAuth.Roles(),
		)
	}
	return nil
}

// RegisterMsgAuth registers the auth to send the msg.
// Each role gets the access control
func (k Keeper) RegisterMsgAuth(msg sdk.Msg, roles ...types.Role) {
	if _, ok := k.AuthMap[sdk.MsgTypeURL(msg)]; ok {
		panic(fmt.Sprintf("msg type or module name %s has already been initialized", sdk.MsgTypeURL(msg)))
	}
	auth := types.AuthDefault
	for _, r := range roles {
		auth = auth | r.Auth()
	}
	k.AuthMap[sdk.MsgTypeURL(msg)] = auth
}

// RegisterModuleAuth registers the auth to send the module related msgs.
// Each role gets the access control
func (k *Keeper) RegisterModuleAuth(module string, roles ...types.Role) {
	if _, ok := k.AuthMap[module]; ok {
		panic(fmt.Sprintf("msg type or module name %s has already been initialized", module))
	}
	auth := types.AuthDefault
	for _, r := range roles {
		auth = auth | r.Auth()
	}
	k.AuthMap[module] = auth
}
