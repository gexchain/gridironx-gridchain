package client

import (
	"github.com/gorilla/mux"

	"github.com/gridironx/gridchain/libs/cosmos-sdk/client/context"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/client/rpc"
)

// Register routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	rpc.RegisterRPCRoutes(cliCtx, r)
}
