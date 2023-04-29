//go:build ignore
// +build ignore

package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenesisState(t *testing.T) {
	tests := []struct {
		pools     FarmPools
		lockInfos []LockInfo
		histories []PoolHistoricalRewardsRecord
		currents  []PoolCurrentRewardsRecord
		whiteList PoolNameList
		params    Params
		err       error
	}{
		{
			pools:     FarmPools{FarmPool{}, FarmPool{}},
			lockInfos: []LockInfo{{}, {}},
			histories: []PoolHistoricalRewardsRecord{
				{Rewards: PoolHistoricalRewards{ReferenceCount: 2}},
				{Rewards: PoolHistoricalRewards{ReferenceCount: 2}},
			},
			currents: []PoolCurrentRewardsRecord{{}, {}},
			err:      nil,
		},
		{
			pools:     FarmPools{FarmPool{}, FarmPool{}},
			lockInfos: []LockInfo{{}, {}},
			histories: []PoolHistoricalRewardsRecord{
				{Rewards: PoolHistoricalRewards{ReferenceCount: 2}},
				{Rewards: PoolHistoricalRewards{ReferenceCount: 2}},
			},
			currents: []PoolCurrentRewardsRecord{{}},
			err:      errors.New(""),
		},
		{
			pools:     FarmPools{FarmPool{}, FarmPool{}},
			lockInfos: []LockInfo{{}, {}},
			histories: []PoolHistoricalRewardsRecord{
				{Rewards: PoolHistoricalRewards{ReferenceCount: 1}},
				{Rewards: PoolHistoricalRewards{ReferenceCount: 2}},
			},
			currents: []PoolCurrentRewardsRecord{{}, {}},
			err:      errors.New(""),
		},
	}

	for _, test := range tests {
		genesis := NewGenesisState(
			test.pools, test.lockInfos, test.histories, test.currents, test.whiteList, test.params,
		)
		if test.err != nil {
			require.Error(t, ValidateGenesis(genesis))
		} else {
			require.NoError(t, ValidateGenesis(genesis))
		}
	}
}
