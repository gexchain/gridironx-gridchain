package types

import (
	"fmt"
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	sdkerrors "github.com/gridironx/gridchain/libs/cosmos-sdk/types/errors"
)

// NOTE: We can't use 1 since that error code is reserved for internal errors.
const (
	DefaultCodespace string = ModuleName
)

var (
	// ErrInvalidState returns an error resulting from an invalid Storage State.
	ErrEmptyTreasures = sdkerrors.Register(ModuleName, 2, "treasures is not empty")

	ErrDuplicatedTreasure      = sdkerrors.Register(ModuleName, 3, "treasures can not be duplicate")
	ErrUnexpectedProposalType  = sdkerrors.Register(ModuleName, 4, "unsupported proposal type of mint module")
	ErrProposerMustBeValidator = sdkerrors.Register(ModuleName, 5, "the proposal of proposer must be validator")
	ErrNotReachedVenus5Height  = sdkerrors.Register(ModuleName, 6, "venus5 block height has not been reached")
	ErrNextBlockUpdateTooLate  = sdkerrors.Register(ModuleName, 7, "the next block to update is too late")
	ErrCodeInvalidHeight       = sdkerrors.Register(ModuleName, 8, "height must be greater than current block")
)

// ErrTreasuresInternal returns an error when the length of address list in the proposal is larger than the max limitation
func ErrTreasuresInternal(err error) sdk.EnvelopedErr {
	return sdk.EnvelopedErr{
		Err: sdkerrors.New(
			DefaultParamspace,
			11,
			fmt.Sprintf("treasures error:%s", err.Error()))}
}
