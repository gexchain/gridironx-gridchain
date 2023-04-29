package vmbridge

import (
	"github.com/gridironx/gridchain/x/vmbridge/keeper"
	"github.com/gridironx/gridchain/x/vmbridge/types"
)

var (
	RegisterMsgServer         = types.RegisterMsgServer
	NewMsgServerImpl          = keeper.NewMsgServerImpl
	NewSendToWasmEventHandler = keeper.NewSendToWasmEventHandler
	NewCallToWasmEventHandler = keeper.NewCallToWasmEventHandler
	RegisterSendToEvmEncoder  = keeper.RegisterSendToEvmEncoder
	NewKeeper                 = keeper.NewKeeper
	RegisterInterface         = types.RegisterInterface
)

type (
	MsgSendToEvm = types.MsgSendToEvm
	Keeper       = keeper.Keeper
)
