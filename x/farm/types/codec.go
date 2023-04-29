package types

import (
	"github.com/gridironx/gridchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreatePool{}, "gridchain/farm/MsgCreatePool", nil)
	cdc.RegisterConcrete(MsgDestroyPool{}, "gridchain/farm/MsgDestroyPool", nil)
	cdc.RegisterConcrete(MsgLock{}, "gridchain/farm/MsgLock", nil)
	cdc.RegisterConcrete(MsgUnlock{}, "gridchain/farm/MsgUnlock", nil)
	cdc.RegisterConcrete(MsgClaim{}, "gridchain/farm/MsgClaim", nil)
	cdc.RegisterConcrete(MsgProvide{}, "gridchain/farm/MsgProvide", nil)
	cdc.RegisterConcrete(ManageWhiteListProposal{}, "gridchain/farm/ManageWhiteListProposal", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
