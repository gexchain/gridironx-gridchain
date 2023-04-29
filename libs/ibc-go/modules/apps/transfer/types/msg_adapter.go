package types

import (
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	sdkerrors "github.com/gridironx/gridchain/libs/cosmos-sdk/types/errors"
)

//for denom convert wei to fury and reject fury direct
func (m *MsgTransfer) RulesFilter() (sdk.Msg, error) {
	if m.Token.Denom == sdk.DefaultBondDenom {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "ibc MsgTransfer not support fury denom")
	}
	return m, nil
}
