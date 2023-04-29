package keeper

import (
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	"github.com/gridironx/gridchain/x/order/types"
)

// nolint
func (k Keeper) IsProductLocked(ctx sdk.Context, product string) bool {
	return k.dexKeeper.IsTokenPairLocked(ctx, product)
}

// nolint
func (k Keeper) SetProductLock(ctx sdk.Context, product string, lock *types.ProductLock) {
	k.dexKeeper.LockTokenPair(ctx, product, lock)
}

// nolint
func (k Keeper) UnlockProduct(ctx sdk.Context, product string) {
	k.dexKeeper.UnlockTokenPair(ctx, product)
}

// nolint
func (k Keeper) AnyProductLocked(ctx sdk.Context) bool {
	return k.dexKeeper.IsAnyProductLocked(ctx)
}
