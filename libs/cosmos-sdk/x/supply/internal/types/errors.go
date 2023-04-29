package types

import (
	sdkerrors "github.com/gridironx/gridchain/libs/cosmos-sdk/types/errors"
)

var (
	ErrInvalidDeflation = sdkerrors.Register(ModuleName, 1, "failed. the deflation is larger than the current supply")
)
