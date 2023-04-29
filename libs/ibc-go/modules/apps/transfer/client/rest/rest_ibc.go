package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gridironx/gridchain/libs/cosmos-sdk/client/context"
	"github.com/gridironx/gridchain/libs/ibc-go/modules/apps/transfer/types"
)

func RegisterOriginRPCRoutersForGRPC(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/ibc/apps/transfer/v1/denom_traces",
		denomTracesHandlerFn(cliCtx),
	).Methods("GET")
}

func denomTracesHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return queryDenomTraces(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDenomTraces))
}
