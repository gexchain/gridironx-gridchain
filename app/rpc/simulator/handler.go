package simulator

import (
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
)

type Simulator interface {
	Simulate([]sdk.Msg) (*sdk.Result, error)
	Context() *sdk.Context
	Release()
}

var NewWasmSimulator func() Simulator
