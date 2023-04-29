package client

import (
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/mint/client/cli"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/mint/client/rest"
	govcli "github.com/gridironx/gridchain/x/gov/client"
)

var (
	ManageTreasuresProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdManageTreasuresProposal,
		rest.ManageTreasuresProposalRESTHandler,
	)
	ModifyNextBlockUpdateProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdModifyNextBlockUpdateProposal,
		rest.ModifyNextBlockUpdateProposalRESTHandler,
	)
)
