package keeper

import sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"

func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	return k.isBound(ctx, portID)
}
