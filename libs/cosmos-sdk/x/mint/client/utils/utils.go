package utils

import (
	"github.com/gridironx/gridchain/libs/cosmos-sdk/codec"
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/mint/internal/types"
	"io/ioutil"
)

// ManageTreasuresProposalJSON defines a ManageTreasureProposal with a deposit used to parse
// manage treasures proposals from a JSON file.
type ManageTreasuresProposalJSON struct {
	Title       string           `json:"title" yaml:"title"`
	Description string           `json:"description" yaml:"description"`
	Treasures   []types.Treasure `json:"treasures" yaml:"treasures"`
	IsAdded     bool             `json:"is_added" yaml:"is_added"`
	Deposit     sdk.SysCoins     `json:"deposit" yaml:"deposit"`
}

// ParseManageTreasuresProposalJSON parses json from proposal file to ManageTreasuresProposalJSON struct
func ParseManageTreasuresProposalJSON(cdc *codec.Codec, proposalFilePath string) (
	proposal ManageTreasuresProposalJSON, err error) {
	contents, err := ioutil.ReadFile(proposalFilePath)
	if err != nil {
		return
	}

	cdc.MustUnmarshalJSON(contents, &proposal)
	return
}

// ModifyNextBlockUpdateProposalJSON defines a ModifyNextBlockUpdateProposal with a deposit used to parse
// manage treasures proposals from a JSON file.
type ModifyNextBlockUpdateProposalJSON struct {
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.SysCoins `json:"deposit" yaml:"deposit"`
	BlockNum    uint64       `json:"block_num" yaml:"block_num"`
}

// ParseModifyNextBlockUpdateProposalJSON parses json from proposal file to ModifyNextBlockUpdateProposalJSON struct
func ParseModifyNextBlockUpdateProposalJSON(cdc *codec.Codec, proposalFilePath string) (
	proposal ModifyNextBlockUpdateProposalJSON, err error) {
	contents, err := ioutil.ReadFile(proposalFilePath)
	if err != nil {
		return
	}

	cdc.MustUnmarshalJSON(contents, &proposal)
	return
}
