package keeper

import (
	"mini_ibc/x/miniibc/types"
)

var _ types.QueryServer = Keeper{}
