package trace

const (
	//----- DeliverTx
	DeliverTx = "DeliverTx"
	TxDecoder = "TxDecoder"

	//----- RunTx details
	ValTxMsgs   = "valTxMsgs"
	RunAnte     = "RunAnte"
	RunMsg      = "RunMsg"
	Refund      = "refund"
	EvmHandler  = "EvmHandler"

	//------ RunAnte details
	CacheTxContext  = "cacheTxContext"
	AnteChain       = "AnteChain"
	AnteOther       = "AnteOther"
	CacheStoreWrite = "cacheStoreWrite"
	//----- RunMsgs details

	//----- handler details
	ParseChainID = "ParseChainID"
	VerifySig    = "VerifySig"
	Txhash       = "txhash"
	SaveTx       = "SaveTx"
	TransitionDb = "TransitionDb"
	Bloomfilter  = "Bloomfilter"
	EmitEvents   = "EmitEvents"
	HandlerDefer = "handler_defer"
)


const (
	GasUsed          = "GasUsed"
	Produce          = "Produce"
	RunTx            = "RunTx"
	Height           = "Height"
	Tx               = "Tx"
	BlockSize        = "BlockSize"
	Elapsed          = "Elapsed"
	CommitRound      = "CommitRound"
	Round            = "Round"
	Evm              = "Evm"
	Iavl             = "Iavl"
	FlatKV           = "FlatKV"
	WtxRatio         = "WtxRatio"
	SigCacheRatio    = "SigCacheRatio"
	DeliverTxs       = "DeliverTxs"
	EvmHandlerDetail = "EvmHandlerDetail"
	RunAnteDetail    = "RunAnteDetail"
	AnteChainDetail  = "AnteChainDetail"

	Delta = "Delta"
	InvalidTxs = "InvalidTxs"

	Abci       = "abci"
	SaveResp   = "saveResp"
	Persist    = "persist"
	SaveState  = "saveState"
	Evpool     = "evpool"
	FireEvents = "fireEvents"
	ApplyBlock = "ApplyBlock"
	Consensus  = "Consensus"

	MempoolCheckTxCnt = "checkTxCnt"
	MempoolTxsCnt     = "mempoolTxsCnt"

	Prerun = "Prerun"
)

const (
	READ         = 1
	WRITE        = 2
	EVMALL       = 3
	UNKNOWN_TYPE = 4
	EVM_FORMAT   = "read<%dms>, write<%dms>, execute<%dms>"
	EVMCORE      = "evmcore"
)

var (
	STATEDB_WRITE = []string{"AddBalance", "SubBalance", "SetNonce", "SetState", "SetCode", "AddLog",
		"AddPreimage", "AddRefund", "SubRefund", "AddAddressToAccessList", "AddSlotToAccessList",
		"PrepareAccessList", "AddressInAccessList", "Suicide", "CreateAccount", "ForEachStorage"}

	STATEDB_READ  = []string{"SlotInAccessList", "GetBalance", "GetNonce", "GetCode", "GetCodeSize",
		"GetCodeHash", "GetState", "GetCommittedState", "GetRefund",
		"HasSuicided", "Snapshot", "RevertToSnapshot", "Empty", "Exist"}

	EVM_OPER      = []string{EVMCORE}
	dbOper        *DbRecord
)