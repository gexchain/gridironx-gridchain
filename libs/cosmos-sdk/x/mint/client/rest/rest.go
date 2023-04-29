package rest

import (
	"github.com/gorilla/mux"
	govRest "github.com/gridironx/gridchain/x/gov/client/rest"

	"github.com/gridironx/gridchain/libs/cosmos-sdk/client/context"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
}

// ManageTreasuresProposalRESTHandler defines mint proposal handler
func ManageTreasuresProposalRESTHandler(context.CLIContext) govRest.ProposalRESTHandler {
	return govRest.ProposalRESTHandler{}
}

// ModifyNextBlockUpdateProposalRESTHandler defines mint proposal handler
func ModifyNextBlockUpdateProposalRESTHandler(context.CLIContext) govRest.ProposalRESTHandler {
	return govRest.ProposalRESTHandler{}
}
