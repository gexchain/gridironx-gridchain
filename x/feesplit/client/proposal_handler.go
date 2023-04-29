package client

import (
	"github.com/gridironx/gridchain/x/feesplit/client/cli"
	"github.com/gridironx/gridchain/x/feesplit/client/rest"
	govcli "github.com/gridironx/gridchain/x/gov/client"
)

var (
	// FeeSplitSharesProposalHandler alias gov NewProposalHandler
	FeeSplitSharesProposalHandler = govcli.NewProposalHandler(
		cli.GetCmdFeeSplitSharesProposal,
		rest.FeeSplitSharesProposalRESTHandler,
	)
)
