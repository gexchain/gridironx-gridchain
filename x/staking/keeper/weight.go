package keeper

import (
	"fmt"
	"math"

	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	types2 "github.com/gridironx/gridchain/libs/tendermint/types"
	"github.com/gridironx/gridchain/x/staking/types"
)

const (
	// UTC Time: 2000/1/1 00:00:00
	blockTimestampEpoch = int64(946684800)
	secondsPerWeek      = int64(60 * 60 * 24 * 7)
	weeksPerYear        = float64(52)
	fixedWeight         = int64(11700000) // The weight of 1 fury, calculated by calculateWeight before venus6. (nowTime=2023-06-01 00:00:00 GMT+0)
)

func calculateWeight(nowTime int64, tokens sdk.Dec, height int64) (shares types.Shares, sdkErr error) {
	if types2.HigherThanVenus6(height) {
		shares = tokens.MulInt64(fixedWeight)
		return
	}

	nowWeek := (nowTime - blockTimestampEpoch) / secondsPerWeek
	rate := float64(nowWeek) / weeksPerYear
	weight := math.Pow(float64(2), rate)

	precision := fmt.Sprintf("%d", sdk.Precision)

	weightByDec, sdkErr := sdk.NewDecFromStr(fmt.Sprintf("%."+precision+"f", weight))
	if sdkErr == nil {
		shares = tokens.Mul(weightByDec)
	}
	return
}

func SimulateWeight(nowTime int64, tokens sdk.Dec, height int64) (votes types.Shares, sdkErr error) {
	return calculateWeight(nowTime, tokens, height)
}
