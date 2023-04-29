package client

import (
	govclient "github.com/gridironx/gridchain/libs/cosmos-sdk/x/gov/client"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/upgrade/client/cli"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
