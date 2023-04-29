package app

import (
	"runtime"
	"time"

	appconfig "github.com/gridironx/gridchain/app/config"
	"github.com/gridironx/gridchain/libs/system/trace"
	abci "github.com/gridironx/gridchain/libs/tendermint/abci/types"
	"github.com/gridironx/gridchain/x/wasm/watcher"
)

// BeginBlock implements the Application interface
func (app *GRIDIronxChainApp) BeginBlock(req abci.RequestBeginBlock) (res abci.ResponseBeginBlock) {
	trace.OnAppBeginBlockEnter(app.LastBlockHeight() + 1)
	app.EvmKeeper.Watcher.DelayEraseKey()
	return app.BaseApp.BeginBlock(req)
}

func (app *GRIDIronxChainApp) DeliverTx(req abci.RequestDeliverTx) (res abci.ResponseDeliverTx) {

	trace.OnAppDeliverTxEnter()

	resp := app.BaseApp.DeliverTx(req)

	return resp
}

func (app *GRIDIronxChainApp) PreDeliverRealTx(req []byte) (res abci.TxEssentials) {
	return app.BaseApp.PreDeliverRealTx(req)
}

func (app *GRIDIronxChainApp) DeliverRealTx(req abci.TxEssentials) (res abci.ResponseDeliverTx) {
	trace.OnAppDeliverTxEnter()
	resp := app.BaseApp.DeliverRealTx(req)
	app.EvmKeeper.Watcher.RecordTxAndFailedReceipt(req, &resp, app.GetTxDecoder())

	return resp
}

// EndBlock implements the Application interface
func (app *GRIDIronxChainApp) EndBlock(req abci.RequestEndBlock) (res abci.ResponseEndBlock) {
	return app.BaseApp.EndBlock(req)
}

// Commit implements the Application interface
func (app *GRIDIronxChainApp) Commit(req abci.RequestCommit) abci.ResponseCommit {
	if gcInterval := appconfig.GetOecConfig().GetGcInterval(); gcInterval > 0 {
		if (app.BaseApp.LastBlockHeight()+1)%int64(gcInterval) == 0 {
			startTime := time.Now()
			runtime.GC()
			elapsed := time.Now().Sub(startTime).Milliseconds()
			app.Logger().Info("force gc for debug", "height", app.BaseApp.LastBlockHeight()+1,
				"elapsed(ms)", elapsed)
		}
	}
	//defer trace.GetTraceSummary().Dump()
	defer trace.OnCommitDone()

	// reload upgrade info for upgrade proposal
	app.setupUpgradeModules(true)
	tasks := app.heightTasks[app.BaseApp.LastBlockHeight()+1]
	if tasks != nil {
		ctx := app.BaseApp.GetDeliverStateCtx()
		for _, t := range *tasks {
			if err := t.Execute(ctx); nil != err {
				panic("bad things")
			}
		}
	}
	res := app.BaseApp.Commit(req)

	// we call watch#Commit here ,because
	// 1. this round commit a valid block
	// 2. before commit the block,State#updateToState hasent not called yet,so the proposalBlockPart is not nil which means we wont
	// 	  call the prerun during commit step(edge case)
	app.EvmKeeper.Watcher.Commit()
	watcher.Commit()

	return res
}
