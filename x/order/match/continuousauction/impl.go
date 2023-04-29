package continuousauction

import (
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"

	"github.com/gridironx/gridchain/x/order/keeper"
)

// nolint
type CaEngine struct {
}

// nolint
func (e *CaEngine) Run(ctx sdk.Context, keeper keeper.Keeper) {
	// TODO
}
