package order

import (
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"

	"github.com/gridironx/gridchain/x/common/perf"
	"github.com/gridironx/gridchain/x/order/keeper"
	"github.com/gridironx/gridchain/x/order/types"
	//"github.com/gridironx/gridchain/x/common/version"
)

// BeginBlocker runs the logic of BeginBlocker with version 0.
// BeginBlocker resets keeper cache.
func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	seq := perf.GetPerf().OnBeginBlockEnter(ctx, types.ModuleName)
	defer perf.GetPerf().OnBeginBlockExit(ctx, types.ModuleName, seq)

	keeper.ResetCache(ctx)
}
