// nolint
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/gridironx/gridchain/x/distribution/types
// ALIASGEN: github.com/gridironx/gridchain/x/distribution/client
package distribution

import (
	"github.com/gridironx/gridchain/x/distribution/client"
	"github.com/gridironx/gridchain/x/distribution/types"
)

var (
	NewMsgWithdrawDelegatorReward          = types.NewMsgWithdrawDelegatorReward
	CommunityPoolSpendProposalHandler      = client.CommunityPoolSpendProposalHandler
	ChangeDistributionTypeProposalHandler  = client.ChangeDistributionTypeProposalHandler
	WithdrawRewardEnabledProposalHandler   = client.WithdrawRewardEnabledProposalHandler
	RewardTruncatePrecisionProposalHandler = client.RewardTruncatePrecisionProposalHandler
	NewMsgWithdrawDelegatorAllRewards      = types.NewMsgWithdrawDelegatorAllRewards
)
