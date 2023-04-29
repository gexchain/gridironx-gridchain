package infura

import (
	"github.com/gridironx/gridchain/libs/tendermint/libs/log"
	"github.com/gridironx/gridchain/x/common/monitor"
	evm "github.com/gridironx/gridchain/x/evm/watcher"
)

// nolint
type Keeper struct {
	metric *monitor.StreamMetrics
	stream *Stream
}

// nolint
func NewKeeper(evmKeeper EvmKeeper, logger log.Logger, metrics *monitor.StreamMetrics) Keeper {
	logger = logger.With("module", "infura")
	k := Keeper{
		metric: metrics,
		stream: NewStream(logger),
	}
	if k.stream.enable {
		evmKeeper.SetObserverKeeper(k)
	}
	return k
}

func (k Keeper) OnSaveTransactionReceipt(tr evm.TransactionReceipt) {
	k.stream.cache.AddTransactionReceipt(tr)
}

func (k Keeper) OnSaveBlock(b evm.Block) {
	k.stream.cache.AddBlock(b)
}

func (k Keeper) OnSaveTransaction(t evm.Transaction) {
	k.stream.cache.AddTransaction(t)
}

func (k Keeper) OnSaveContractCode(address string, code []byte) {
	k.stream.cache.AddContractCode(address, code)
}
