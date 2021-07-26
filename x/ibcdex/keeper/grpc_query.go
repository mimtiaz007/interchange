package keeper

import (
	"github.com/mimtiaz007/interchange/x/ibcdex/types"
)

var _ types.QueryServer = Keeper{}
