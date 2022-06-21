package keeper

import (
	"github.com/goldnet/chain/x/identity/types"
)

var _ types.QueryServer = Keeper{}
