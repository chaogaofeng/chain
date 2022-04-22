package keeper

import (
	"github.com/glodnet/chain/x/token/types"
)

var _ types.QueryServer = Keeper{}
