package appstatus

import (
	"fmt"
	"math"
	"path/filepath"

	bam "github.com/gridironx/gridchain/libs/cosmos-sdk/baseapp"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/client/flags"
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/auth"
	capabilitytypes "github.com/gridironx/gridchain/libs/cosmos-sdk/x/capability/types"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/mint"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/params"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/supply"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/upgrade"
	"github.com/gridironx/gridchain/libs/iavl"
	ibctransfertypes "github.com/gridironx/gridchain/libs/ibc-go/modules/apps/transfer/types"
	ibchost "github.com/gridironx/gridchain/libs/ibc-go/modules/core/24-host"
	dbm "github.com/gridironx/gridchain/libs/tm-db"
	"github.com/gridironx/gridchain/x/ammswap"
	dex "github.com/gridironx/gridchain/x/dex/types"
	distr "github.com/gridironx/gridchain/x/distribution"
	"github.com/gridironx/gridchain/x/erc20"
	"github.com/gridironx/gridchain/x/evidence"
	"github.com/gridironx/gridchain/x/evm"
	"github.com/gridironx/gridchain/x/farm"
	"github.com/gridironx/gridchain/x/feesplit"
	"github.com/gridironx/gridchain/x/gov"
	"github.com/gridironx/gridchain/x/order"
	"github.com/gridironx/gridchain/x/slashing"
	staking "github.com/gridironx/gridchain/x/staking/types"
	token "github.com/gridironx/gridchain/x/token/types"
	"github.com/spf13/viper"
)

const (
	applicationDB = "application"
	dbFolder      = "data"
)

func GetAllStoreKeys() []string {
	return []string{
		bam.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, mint.StoreKey, distr.StoreKey, slashing.StoreKey,
		gov.StoreKey, params.StoreKey, upgrade.StoreKey, evidence.StoreKey,
		evm.StoreKey, token.StoreKey, token.KeyLock, dex.StoreKey, dex.TokenPairStoreKey,
		order.OrderStoreKey, ammswap.StoreKey, farm.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		ibchost.StoreKey,
		erc20.StoreKey,
		// mpt.StoreKey,
		// wasm.StoreKey,
		feesplit.StoreKey,
	}
}

func IsFastStorageStrategy() bool {
	return checkFastStorageStrategy(GetAllStoreKeys())
}

func checkFastStorageStrategy(storeKeys []string) bool {
	home := viper.GetString(flags.FlagHome)
	dataDir := filepath.Join(home, dbFolder)
	db, err := sdk.NewDB(applicationDB, dataDir)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, v := range storeKeys {
		if !isFss(db, v) {
			return false
		}
	}

	return true
}

func isFss(db dbm.DB, storeKey string) bool {
	prefix := fmt.Sprintf("s/k:%s/", storeKey)
	prefixDB := dbm.NewPrefixDB(db, []byte(prefix))

	return iavl.IsFastStorageStrategy(prefixDB)
}

func GetFastStorageVersion() int64 {
	home := viper.GetString(flags.FlagHome)
	dataDir := filepath.Join(home, dbFolder)
	db, err := sdk.NewDB(applicationDB, dataDir)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	storeKeys := GetAllStoreKeys()
	var ret int64 = math.MaxInt64
	for _, v := range storeKeys {
		version := getVersion(db, v)
		if version < ret {
			ret = version
		}
	}

	return ret
}

func getVersion(db dbm.DB, storeKey string) int64 {
	prefix := fmt.Sprintf("s/k:%s/", storeKey)
	prefixDB := dbm.NewPrefixDB(db, []byte(prefix))

	version, _ := iavl.GetFastStorageVersion(prefixDB)

	return version
}
