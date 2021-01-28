package keeper

import (
	"github.com/faddat/blogdemo/x/blogdemo/types"
)

var _ types.QueryServer = Keeper{}
