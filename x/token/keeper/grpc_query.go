package keeper

import (
	"github.com/goldnet/chain/x/token/types"
)

var _ types.QueryServer = Keeper{}
