package keeper

import (
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	govtypes "github.com/gridironx/gridchain/x/gov/types"
)

// GovKeeper defines the expected gov Keeper
type GovKeeper interface {
	GetDepositParams(ctx sdk.Context) govtypes.DepositParams
	GetVotingParams(ctx sdk.Context) govtypes.VotingParams
}
