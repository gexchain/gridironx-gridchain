package types

import (
	sdkerrors "github.com/gridironx/gridchain/libs/cosmos-sdk/types/errors"
)

// Localhost sentinel errors
var (
	ErrConsensusStatesNotStored = sdkerrors.Register(SubModuleName, 2, "localhost does not store consensus states")
)
