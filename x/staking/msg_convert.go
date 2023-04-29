package staking

import (
	"encoding/json"
	"errors"

	"github.com/gridironx/gridchain/libs/cosmos-sdk/baseapp"
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	tmtypes "github.com/gridironx/gridchain/libs/tendermint/types"
	"github.com/gridironx/gridchain/x/common"
	"github.com/gridironx/gridchain/x/staking/types"
)

var (
	ErrCheckSignerFail = errors.New("check signer fail")
)

func init() {
	RegisterConvert()
}

func RegisterConvert() {
	enableHeight := tmtypes.GetVenus3Height()
	baseapp.RegisterCmHandle("gridchain/staking/MsgDeposit", baseapp.NewCMHandle(ConvertDepositMsg, enableHeight))
	baseapp.RegisterCmHandle("gridchain/staking/MsgWithdraw", baseapp.NewCMHandle(ConvertWithdrawMsg, enableHeight))
	baseapp.RegisterCmHandle("gridchain/staking/MsgAddShares", baseapp.NewCMHandle(ConvertAddSharesMsg, enableHeight))
}

func ConvertDepositMsg(data []byte, signers []sdk.AccAddress) (sdk.Msg, error) {
	newMsg := types.MsgDeposit{}
	err := json.Unmarshal(data, &newMsg)
	if err != nil {
		return nil, err
	}
	err = newMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	if ok := common.CheckSignerAddress(signers, newMsg.GetSigners()); !ok {
		return nil, ErrCheckSignerFail
	}
	return newMsg, nil
}

func ConvertWithdrawMsg(data []byte, signers []sdk.AccAddress) (sdk.Msg, error) {
	newMsg := types.MsgWithdraw{}
	err := json.Unmarshal(data, &newMsg)
	if err != nil {
		return nil, err
	}
	err = newMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	if ok := common.CheckSignerAddress(signers, newMsg.GetSigners()); !ok {
		return nil, ErrCheckSignerFail
	}
	return newMsg, nil
}

func ConvertAddSharesMsg(data []byte, signers []sdk.AccAddress) (sdk.Msg, error) {
	newMsg := types.MsgAddShares{}
	err := json.Unmarshal(data, &newMsg)
	if err != nil {
		return nil, err
	}
	err = newMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	if ok := common.CheckSignerAddress(signers, newMsg.GetSigners()); !ok {
		return nil, ErrCheckSignerFail
	}
	return newMsg, nil
}
