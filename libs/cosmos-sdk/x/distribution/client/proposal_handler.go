package client

import (
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/distribution/client/cli"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/distribution/client/rest"
	govclient "github.com/gridironx/gridchain/libs/cosmos-sdk/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
