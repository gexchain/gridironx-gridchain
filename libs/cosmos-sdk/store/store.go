package store

import (
	dbm "github.com/gridironx/gridchain/libs/tm-db"

	"github.com/gridironx/gridchain/libs/cosmos-sdk/store/cache"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/store/rootmulti"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
