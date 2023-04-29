package keeper

import (
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"

	"github.com/gridironx/gridchain/x/distribution/types"
	"github.com/gridironx/gridchain/x/staking/exported"
)

// initialize rewards for a new validator
func (k Keeper) initializeValidator(ctx sdk.Context, val exported.ValidatorI) {
	if k.CheckDistributionProposalValid(ctx) {
		k.initializeValidatorDistrProposal(ctx, val)
		return
	}

	// set accumulated commissions
	k.SetValidatorAccumulatedCommission(ctx, val.GetOperator(), types.InitialValidatorAccumulatedCommission())
}
