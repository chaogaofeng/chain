package keeper

import (
	"github.com/goldnet/chain/x/permission/types"
)

var _ types.QueryServer = Keeper{}
