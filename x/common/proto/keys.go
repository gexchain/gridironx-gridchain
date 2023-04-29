package proto

import (
	"github.com/gridironx/gridchain/libs/cosmos-sdk/codec"
)

// keys
var (
	upgradeConfigKey     = []byte("upgrade_config")
	currentVersionKey    = []byte("current_version")
	lastFailedVersionKey = []byte("last_failed_version")
	cdc                  = codec.New()
)
