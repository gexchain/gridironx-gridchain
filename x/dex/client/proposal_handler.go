package client

import (
	"github.com/gridironx/gridchain/x/dex/client/cli"
	"github.com/gridironx/gridchain/x/dex/client/rest"
	govclient "github.com/gridironx/gridchain/x/gov/client"
)

// param change proposal handler
var (
	// DelistProposalHandler alias gov NewProposalHandler
	DelistProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitDelistProposal, rest.DelistProposalRESTHandler)
)
