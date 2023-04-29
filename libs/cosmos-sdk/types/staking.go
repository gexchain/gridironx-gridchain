package types

import (
	"math/big"
)

// staking constants
const (

	// default bond denomination
	DefaultBondDenom = "fury"

	// Delay, in blocks, between when validator updates are returned to the
	// consensus-engine and when they are applied. For example, if
	// ValidatorUpdateDelay is set to X, and if a validator set update is
	// returned with new validators at the end of block 10, then the new
	// validators are expected to sign blocks beginning at block 11+X.
	//
	// This value is constant as this should not change without a hard fork.
	// For Tendermint this should be set to 1 block, for more details see:
	// https://tendermint.com/docs/spec/abci/apps.html#endblock
	ValidatorUpdateDelay int64 = 1
)

// PowerReduction is the amount of staking tokens required for 1 unit of consensus-engine power
var PowerReduction = NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(Precision), nil))

// TokensToConsensusPower - convert input tokens to potential consensus-engine power
func TokensToConsensusPower(tokens Int) int64 {
	return (tokens.Quo(PowerReduction)).Int64()
}

// TokensFromConsensusPower - convert input power to tokens
func TokensFromConsensusPower(power int64) Int {
	return NewInt(power).Mul(PowerReduction)
}

// BondStatus is the status of a validator
type BondStatus byte

// staking constants
const (
	Unbonded  BondStatus = 0x00
	Unbonding BondStatus = 0x01
	Bonded    BondStatus = 0x02

	BondStatusUnbonded  = "Unbonded"
	BondStatusUnbonding = "Unbonding"
	BondStatusBonded    = "Bonded"

	CM45BondStatusUnbonded  = "BOND_STATUS_UNBONDED"
	CM45BondStatusUnbonding = "BOND_STATUS_UNBONDING"
	CM45BondStatusBonded    = "BOND_STATUS_BONDED"
)

// Equal compares two BondStatus instances
func (b BondStatus) Equal(b2 BondStatus) bool {
	return byte(b) == byte(b2)
}

// String implements the Stringer interface for BondStatus.
func (b BondStatus) String() string {
	switch b {
	case 0x00:
		return BondStatusUnbonded
	case 0x01:
		return BondStatusUnbonding
	case 0x02:
		return BondStatusBonded
	default:
		panic("invalid bond status")
	}
}

// CM45String is used to compatible with cosmos v0.45.1
func (b BondStatus) CM45String() string {
	switch b {
	case 0x00:
		return CM45BondStatusUnbonded
	case 0x01:
		return CM45BondStatusUnbonding
	case 0x02:
		return CM45BondStatusBonded
	default:
		panic("invalid bond status")
	}
}
