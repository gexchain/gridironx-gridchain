package ante

import (
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	tmtypes "github.com/gridironx/gridchain/libs/tendermint/types"
	"github.com/gridironx/gridchain/x/evm/types"
	wasmkeeper "github.com/gridironx/gridchain/x/wasm/keeper"
)

type WrapWasmCountTXDecorator struct {
	ctd       *wasmkeeper.CountTXDecorator
	evmKeeper EVMKeeper
}

// NewWrapWasmCountTXDecorator constructor
func NewWrapWasmCountTXDecorator(ctd *wasmkeeper.CountTXDecorator, evmKeeper EVMKeeper) *WrapWasmCountTXDecorator {
	return &WrapWasmCountTXDecorator{ctd: ctd, evmKeeper: evmKeeper}
}

func (a WrapWasmCountTXDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if tmtypes.HigherThanVenus6(ctx.BlockHeight()) && isE2CTx(a.evmKeeper, &ctx, tx) {
		return a.ctd.AnteHandle(ctx, tx, simulate, next)
	}
	return next(ctx, tx, simulate)
}

func isE2CTx(ek EVMKeeper, ctx *sdk.Context, tx sdk.Tx) bool {
	evmTx, ok := tx.(*types.MsgEthereumTx)
	if !ok {
		return false
	}
	return IsE2CTx(ek, ctx, evmTx)
}
