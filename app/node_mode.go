package app

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/viper"

	"github.com/gridironx/gridchain/app/config"
	appconfig "github.com/gridironx/gridchain/app/config"
	"github.com/gridironx/gridchain/app/rpc/backend"
	"github.com/gridironx/gridchain/app/types"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/client/flags"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/server"
	store "github.com/gridironx/gridchain/libs/cosmos-sdk/store/iavl"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/types/innertx"
	"github.com/gridironx/gridchain/libs/iavl"
	abcitypes "github.com/gridironx/gridchain/libs/tendermint/abci/types"
	"github.com/gridironx/gridchain/libs/tendermint/libs/log"
	"github.com/gridironx/gridchain/libs/tendermint/mempool"
	tmtypes "github.com/gridironx/gridchain/libs/tendermint/types"
	evmtypes "github.com/gridironx/gridchain/x/evm/types"
	"github.com/gridironx/gridchain/x/evm/watcher"
)

func setNodeConfig(ctx *server.Context) error {
	nodeMode := viper.GetString(types.FlagNodeMode)

	ctx.Logger.Info("Starting node", "mode", nodeMode)

	switch types.NodeMode(nodeMode) {
	case types.RpcNode:
		setRpcConfig(ctx)
	case types.ValidatorNode:
		setValidatorConfig(ctx)
	case types.ArchiveNode:
		setArchiveConfig(ctx)
	case types.InnertxNode:
		if !innertx.IsAvailable {
			return errors.New("innertx is not available for innertx node")
		}
		setRpcConfig(ctx)
	default:
		if len(nodeMode) > 0 {
			ctx.Logger.Error(
				fmt.Sprintf("Wrong value (%s) is set for %s, the correct value should be one of %s, %s, and %s",
					nodeMode, types.FlagNodeMode, types.RpcNode, types.ValidatorNode, types.ArchiveNode))
		}
	}
	return nil
}

func setRpcConfig(ctx *server.Context) {
	viper.SetDefault(abcitypes.FlagDisableABCIQueryMutex, true)
	viper.SetDefault(evmtypes.FlagEnableBloomFilter, true)
	viper.SetDefault(watcher.FlagFastQueryLru, 10000)
	viper.SetDefault(watcher.FlagFastQuery, true)
	viper.SetDefault(backend.FlagApiBackendBlockLruCache, 30000)
	viper.SetDefault(backend.FlagApiBackendTxLruCache, 100000)
	viper.SetDefault(iavl.FlagIavlEnableAsyncCommit, true)
	viper.SetDefault(flags.FlagMaxOpenConnections, 20000)
	viper.SetDefault(mempool.FlagEnablePendingPool, true)
	viper.SetDefault(server.FlagCORS, "*")
	ctx.Logger.Info(fmt.Sprintf(
		"Set --%s=%v\n--%s=%v\n--%s=%v\n--%s=%v\n--%s=%v\n--%s=%v\n--%s=%v\n--%s=%v by rpc node mode",
		abcitypes.FlagDisableABCIQueryMutex, true, evmtypes.FlagEnableBloomFilter, true, watcher.FlagFastQueryLru, 10000,
		watcher.FlagFastQuery, true, iavl.FlagIavlEnableAsyncCommit, true,
		flags.FlagMaxOpenConnections, 20000, mempool.FlagEnablePendingPool, true,
		server.FlagCORS, "*"))
}

func setValidatorConfig(ctx *server.Context) {
	viper.SetDefault(abcitypes.FlagDisableABCIQueryMutex, true)
	viper.SetDefault(appconfig.FlagDynamicGpMode, tmtypes.MinimalGpMode)
	viper.SetDefault(iavl.FlagIavlEnableAsyncCommit, true)
	viper.SetDefault(store.FlagIavlCacheSize, 10000000)
	viper.SetDefault(server.FlagPruning, "everything")
	viper.SetDefault(evmtypes.FlagEnableBloomFilter, false)
	viper.SetDefault(watcher.FlagFastQuery, false)
	viper.SetDefault(appconfig.FlagMaxGasUsedPerBlock, 120000000)
	viper.SetDefault(mempool.FlagEnablePendingPool, false)
	viper.SetDefault(config.FlagEnablePGU, true)

	ctx.Logger.Info(fmt.Sprintf("Set --%s=%v\n--%s=%v\n--%s=%v\n--%s=%v\n--%s=%v\n--%s=%v\n--%s=%v\n--%s=%v\n--%s=%v by validator node mode",
		abcitypes.FlagDisableABCIQueryMutex, true, appconfig.FlagDynamicGpMode, tmtypes.MinimalGpMode, iavl.FlagIavlEnableAsyncCommit, true,
		store.FlagIavlCacheSize, 10000000, server.FlagPruning, "everything",
		evmtypes.FlagEnableBloomFilter, false, watcher.FlagFastQuery, false, appconfig.FlagMaxGasUsedPerBlock, 120000000,
		mempool.FlagEnablePendingPool, false))
}

func setArchiveConfig(ctx *server.Context) {
	viper.SetDefault(server.FlagPruning, "nothing")
	viper.SetDefault(abcitypes.FlagDisableABCIQueryMutex, true)
	viper.SetDefault(evmtypes.FlagEnableBloomFilter, true)
	viper.SetDefault(iavl.FlagIavlEnableAsyncCommit, true)
	viper.SetDefault(flags.FlagMaxOpenConnections, 20000)
	viper.SetDefault(server.FlagCORS, "*")
	ctx.Logger.Info(fmt.Sprintf(
		"Set --%s=%v\n--%s=%v\n--%s=%v\n--%s=%v\n--%s=%v\n--%s=%v by archive node mode",
		server.FlagPruning, "nothing", abcitypes.FlagDisableABCIQueryMutex, true, evmtypes.FlagEnableBloomFilter, true,
		iavl.FlagIavlEnableAsyncCommit, true, flags.FlagMaxOpenConnections, 20000,
		server.FlagCORS, "*"))
}

func logStartingFlags(logger log.Logger) {
	msg := "All flags:\n"

	var maxLen int
	kvMap := make(map[string]interface{})
	var keys []string
	for _, key := range viper.AllKeys() {
		if strings.Index(key, "infura.") == 0 {
			continue
		}

		keys = append(keys, key)
		kvMap[key] = viper.Get(key)
		if len(key) > maxLen {
			maxLen = len(key)
		}
	}

	sort.Strings(keys)
	for _, k := range keys {
		msg += fmt.Sprintf("	%-45s= %v\n", k, kvMap[k])
	}

	logger.Info(msg)
}
