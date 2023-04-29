package client

import (
	"github.com/gridironx/gridchain/x/farm/client/cli"
	"github.com/gridironx/gridchain/x/farm/client/rest"
	govcli "github.com/gridironx/gridchain/x/gov/client"
)

var (
	// ManageWhiteListProposalHandler alias gov NewProposalHandler
	ManageWhiteListProposalHandler = govcli.NewProposalHandler(cli.GetCmdManageWhiteListProposal, rest.ManageWhiteListProposalRESTHandler)
)
