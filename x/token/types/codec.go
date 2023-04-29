package types

import (
	"github.com/gridironx/gridchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTokenIssue{}, "gridchain/token/MsgIssue", nil)
	cdc.RegisterConcrete(MsgTokenBurn{}, "gridchain/token/MsgBurn", nil)
	cdc.RegisterConcrete(MsgTokenMint{}, "gridchain/token/MsgMint", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "gridchain/token/MsgMultiTransfer", nil)
	cdc.RegisterConcrete(MsgSend{}, "gridchain/token/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgTransferOwnership{}, "gridchain/token/MsgTransferOwnership", nil)
	cdc.RegisterConcrete(MsgConfirmOwnership{}, "gridchain/token/MsgConfirmOwnership", nil)
	cdc.RegisterConcrete(MsgTokenModify{}, "gridchain/token/MsgModify", nil)

	// for test
	//cdc.RegisterConcrete(MsgTokenDestroy{}, "gridchain/token/MsgDestroy", nil)
}

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
