package staking

import (
	"context"

	"github.com/gridironx/gridchain/x/staking/keeper"

	cosmost "github.com/gridironx/gridchain/libs/cosmos-sdk/store/types"
	"github.com/gridironx/gridchain/x/staking/typesadapter"

	"github.com/gorilla/mux"
	clictx "github.com/gridironx/gridchain/libs/cosmos-sdk/client/context"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/codec"
	anytypes "github.com/gridironx/gridchain/libs/cosmos-sdk/codec/types"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/gridironx/gridchain/libs/cosmos-sdk/types/upgrade"
	params2 "github.com/gridironx/gridchain/libs/cosmos-sdk/x/params"
	"github.com/gridironx/gridchain/x/params"
	"github.com/gridironx/gridchain/x/staking/client/rest"
	"github.com/gridironx/gridchain/x/staking/types"
	_ "github.com/gridironx/gridchain/x/staking/typesadapter"
	"github.com/spf13/cobra"
)

var (
	_ upgrade.UpgradeModule        = AppModule{}
	_ module.AppModuleAdapter      = AppModule{}
	_ module.AppModuleBasicAdapter = AppModuleBasic{}
)

// appmoduleBasic
func (am AppModuleBasic) RegisterRouterForGRPC(cliCtx clictx.CLIContext, r *mux.Router) {
	rest.RegisterOriginRPCRoutersForGRPC(cliCtx, r)
}

func (am AppModuleBasic) RegisterInterfaces(registry anytypes.InterfaceRegistry) {}

func (am AppModuleBasic) RegisterGRPCGatewayRoutes(cliContext clictx.CLIContext, serveMux *runtime.ServeMux) {
	typesadapter.RegisterQueryHandlerClient(context.Background(), serveMux, typesadapter.NewQueryClient(cliContext))
}

func (am AppModuleBasic) GetTxCmdV2(cdc *codec.CodecProxy, reg anytypes.InterfaceRegistry) *cobra.Command {
	return nil
}

func (am AppModuleBasic) GetQueryCmdV2(cdc *codec.CodecProxy, reg anytypes.InterfaceRegistry) *cobra.Command {
	return nil
}

// / appmodule
func (am AppModule) RegisterServices(cfg module.Configurator) {
	typesadapter.RegisterQueryServer(cfg.QueryServer(), keeper.NewGrpcQuerier(am.keeper))
}

func (am AppModule) RegisterTask() upgrade.HeightTask {
	return nil
}

func (am AppModule) UpgradeHeight() int64 {
	return -1
}

func (am AppModule) RegisterParam() params.ParamSet {
	v := types.KeyHistoricalEntriesParams(7)
	return params2.ParamSet(v)
}

func (am AppModule) ModuleName() string {
	return ModuleName
}

func (am AppModule) CommitFilter() *cosmost.StoreFilter {
	return nil
}

func (am AppModule) PruneFilter() *cosmost.StoreFilter {
	return nil
}

func (am AppModule) VersionFilter() *cosmost.VersionFilter {
	return nil
}
