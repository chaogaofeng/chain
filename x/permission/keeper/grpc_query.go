package keeper

import (
	"github.com/glodnet/chain/x/permission/types"
)

var _ types.QueryServer = Keeper{}
