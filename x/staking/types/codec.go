package types

import (
	"github.com/gridironx/gridchain/libs/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types for codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "gridchain/staking/MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgEditValidator{}, "gridchain/staking/MsgEditValidator", nil)
	cdc.RegisterConcrete(MsgEditValidatorCommissionRate{}, "gridchain/staking/MsgEditValidatorCommissionRate", nil)
	cdc.RegisterConcrete(MsgDestroyValidator{}, "gridchain/staking/MsgDestroyValidator", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "gridchain/staking/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "gridchain/staking/MsgWithdraw", nil)
	cdc.RegisterConcrete(MsgAddShares{}, "gridchain/staking/MsgAddShares", nil)
	cdc.RegisterConcrete(MsgRegProxy{}, "gridchain/staking/MsgRegProxy", nil)
	cdc.RegisterConcrete(MsgBindProxy{}, "gridchain/staking/MsgBindProxy", nil)
	cdc.RegisterConcrete(MsgUnbindProxy{}, "gridchain/staking/MsgUnbindProxy", nil)
	cdc.RegisterConcrete(CM45Validator{}, "cosmos-sdk/staking/validator", nil)
}

// ModuleCdc is generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
