package keeper

import (
	"github.com/glodnet/chain/x/identity/types"
)

var _ types.QueryServer = Keeper{}
