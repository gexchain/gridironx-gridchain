package vmbridge

import (
	"github.com/gridironx/gridchain/libs/cosmos-sdk/codec"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/types/module"
	"github.com/gridironx/gridchain/x/vmbridge/keeper"
	"github.com/gridironx/gridchain/x/wasm"
)

func RegisterServices(cfg module.Configurator, keeper keeper.Keeper) {
	RegisterMsgServer(cfg.MsgServer(), NewMsgServerImpl(keeper))
}

func GetWasmOpts(cdc *codec.ProtoCodec) wasm.Option {
	return wasm.WithMessageEncoders(RegisterSendToEvmEncoder(cdc))
}
