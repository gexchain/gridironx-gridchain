maDEP := $(shell command -v dep 2> /dev/null)
SUM := $(shell which shasum)

COMMIT := $(shell git rev-parse HEAD)
CAT := $(if $(filter $(OS),Windows_NT),type,cat)
export GO111MODULE=on

GithubTop=github.com

GO_VERSION=1.17
ROCKSDB_VERSION=6.27.3
IGNORE_CHECK_GO=false
install_rocksdb_version:=$(ROCKSDB_VERSION)


Version=v1.7.0
CosmosSDK=v0.39.2
Tendermint=v0.33.9
Iavl=v0.14.3
Name=gridchain
ServerName=gridchaind
ClientName=gridchaincli
# the height of the 1st block is GenesisHeight+1
GenesisHeight=0
MercuryHeight=1
VenusHeight=1
Venus1Height=1
Venus2Height=0
Venus3Height=1
Venus4Height=0
Venus5Height=0
EarthHeight=0
MarsHeight=0

LINK_STATICALLY = false
cgo_flags=

ifeq ($(IGNORE_CHECK_GO),true)
    GO_VERSION=0
endif

# process linker flags
ifeq ($(VERSION),)
    VERSION = $(COMMIT)
endif

ifeq ($(MAKECMDGOALS),mainnet)
   GenesisHeight=2322600
   MercuryHeight=5150000
   VenusHeight=8200000
   Venus1Height=12988000
   Venus2Height=14738000
   Venus3Height=15277000
   Venus5Height=17849000
   EarthHeight=18735000
	
   WITH_ROCKSDB=true
else ifeq ($(MAKECMDGOALS),testnet)
   GenesisHeight=1121818
   MercuryHeight=5300000
   VenusHeight=8510000
   Venus1Height=12067000
   Venus2Height=14781000
   Venus3Height=15540000
   EarthHeight=17364500
   Venus4Height=17531500
   Venus5Height=18861500

   WITH_ROCKSDB=true
endif

build_tags = netgo

ifeq ($(WITH_ROCKSDB),true)
  CGO_ENABLED=1
  build_tags += rocksdb
  ifeq ($(LINK_STATICALLY),true)
      cgo_flags += CGO_CFLAGS="-I/usr/include/rocksdb"
      cgo_flags += CGO_LDFLAGS="-L/usr/lib -lrocksdb -lstdc++ -lm  -lsnappy -llz4"
  endif
else
  ROCKSDB_VERSION=0
endif

ifeq ($(LINK_STATICALLY),true)
	build_tags += muslc
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

ldflags = -X $(GithubTop)/gridironx/gridchain/libs/cosmos-sdk/version.Version=$(Version) \
	-X $(GithubTop)/gridironx/gridchain/libs/cosmos-sdk/version.Name=$(Name) \
  -X $(GithubTop)/gridironx/gridchain/libs/cosmos-sdk/version.ServerName=$(ServerName) \
  -X $(GithubTop)/gridironx/gridchain/libs/cosmos-sdk/version.ClientName=$(ClientName) \
  -X $(GithubTop)/gridironx/gridchain/libs/cosmos-sdk/version.Commit=$(COMMIT) \
  -X $(GithubTop)/gridironx/gridchain/libs/cosmos-sdk/version.CosmosSDK=$(CosmosSDK) \
  -X $(GithubTop)/gridironx/gridchain/libs/cosmos-sdk/version.Tendermint=$(Tendermint) \
  -X "$(GithubTop)/gridironx/gridchain/libs/cosmos-sdk/version.BuildTags=$(build_tags)" \
  -X $(GithubTop)/gridironx/gridchain/libs/tendermint/types.MILESTONE_GENESIS_HEIGHT=$(GenesisHeight) \
  -X $(GithubTop)/gridironx/gridchain/libs/tendermint/types.MILESTONE_MERCURY_HEIGHT=$(MercuryHeight) \
  -X $(GithubTop)/gridironx/gridchain/libs/tendermint/types.MILESTONE_VENUS_HEIGHT=$(VenusHeight) \
  -X $(GithubTop)/gridironx/gridchain/libs/tendermint/types.MILESTONE_VENUS1_HEIGHT=$(Venus1Height) \
  -X $(GithubTop)/gridironx/gridchain/libs/tendermint/types.MILESTONE_VENUS2_HEIGHT=$(Venus2Height) \
  -X $(GithubTop)/gridironx/gridchain/libs/tendermint/types.MILESTONE_VENUS3_HEIGHT=$(Venus3Height) \
  -X $(GithubTop)/gridironx/gridchain/libs/tendermint/types.MILESTONE_VENUS4_HEIGHT=$(Venus4Height) \
  -X $(GithubTop)/gridironx/gridchain/libs/tendermint/types.MILESTONE_VENUS5_HEIGHT=$(Venus5Height) \
  -X $(GithubTop)/gridironx/gridchain/libs/tendermint/types.MILESTONE_EARTH_HEIGHT=$(EarthHeight) \
  -X $(GithubTop)/gridironx/gridchain/libs/tendermint/types.MILESTONE_MARS_HEIGHT=$(MarsHeight)


ifeq ($(WITH_ROCKSDB),true)
  ldflags += -X github.com/gridironx/gridchain/libs/tendermint/types.DBBackend=rocksdb
endif

ifeq ($(MAKECMDGOALS),testnet)
  ldflags += -X github.com/gridironx/gridchain/libs/cosmos-sdk/server.ChainID=clockend-420
endif

ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif

ifeq ($(GRIDCMALLOC),tcmalloc)
  ldflags += -extldflags "-ltcmalloc_minimal"
endif

ifeq ($(GRIDCMALLOC),jemalloc)
  ldflags += -extldflags "-ljemalloc"
endif

BUILD_FLAGS := -ldflags '$(ldflags)'

ifeq ($(DEBUG),true)
	BUILD_FLAGS += -gcflags "all=-N -l"
endif

all: install

install: gridchain


gridchain: check_version
	$(cgo_flags) go install -v $(BUILD_FLAGS) -tags "$(build_tags)" ./cmd/gridchaind
	$(cgo_flags) go install -v $(BUILD_FLAGS) -tags "$(build_tags)" ./cmd/gridchaincli

check_version:
	@sh $(shell pwd)/libs/check/check-version.sh $(GO_VERSION) $(ROCKSDB_VERSION)

mainnet: gridchain

testnet: gridchain

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./app/...
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./x/backend/...
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./x/common/...
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./x/dex/...
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./x/distribution/...
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./x/genutil/...
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./x/gov/...
#	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./x/order/...
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./x/params/...
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./x/staking/...
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./x/token/...
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./x/upgrade/...

get_vendor_deps:
	@echo "--> Generating vendor directory via dep ensure"
	@rm -rf .vendor-new
	@dep ensure -v -vendor-only

update_vendor_deps:
	@echo "--> Running dep ensure"
	@rm -rf .vendor-new
	@dep ensure -v -update

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download
.PHONY: go-mod-cache

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify
	@go mod tidy

cli:
	go install -v $(BUILD_FLAGS) -tags "$(build_tags)" ./cmd/gridchaincli

server:
	go install -v $(BUILD_FLAGS) -tags "$(build_tags)" ./cmd/gridchaind

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs gofmt -w -s

build:
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -tags "$(build_tags)" -o build/gridchaind.exe ./cmd/gridchaind
	go build $(BUILD_FLAGS) -tags "$(build_tags)" -o build/gridchaincli.exe ./cmd/gridchaincli
else
	go build $(BUILD_FLAGS) -tags "$(build_tags)" -o build/gridchaind ./cmd/gridchaind
	go build $(BUILD_FLAGS) -tags "$(build_tags)" -o build/gridchaincli ./cmd/gridchaincli
endif


test:
	go list ./app/... |xargs go test -count=1
	go list ./x/... |xargs go test -count=1
	go list ./libs/cosmos-sdk/... |xargs go test -count=1 -tags='norace ledger test_ledger_mock'
	go list ./libs/tendermint/... |xargs go test -count=1
	go list ./libs/tm-db/... |xargs go test -count=1
	go list ./libs/iavl/... |xargs go test -count=1
	go list ./libs/ibc-go/... |xargs go test -count=1

testapp:
	go list ./app/... |xargs go test -count=1

testx:
	go list ./x/... |xargs go test -count=1

testcm:
	go list ./libs/cosmos-sdk/... |xargs go test -count=1 -tags='norace ledger test_ledger_mock'

testtm:
	go list ./libs/tendermint/... |xargs go test -count=1 -tags='norace ledger test_ledger_mock'

testibc:
	go list ./libs/ibc-go/... |xargs go test -count=1 -tags='norace ledger test_ledger_mock'


build-linux:
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

build-docker-gridchainnode:
	$(MAKE) -C networks/local

# Run a 4-node testnet locally
localnet-start: localnet-stop
	@if ! [ -f build/node0/gridchaind/config/genesis.json ]; then docker run --rm -v $(CURDIR)/build:/gridchaind:Z gridchain/node testnet --v 4 -o . --starting-ip-address 192.168.10.2 --keyring-backend=test ; fi
	docker-compose up -d

# Stop testnet
localnet-stop:
	docker-compose down

rocksdb:
	@echo "Installing rocksdb..."
	@bash ./libs/rocksdb/install.sh --version v$(install_rocksdb_version)
.PHONY: rocksdb

.PHONY: build

tcmalloc:
	@echo "Installing tcmalloc..."
	@bash ./libs/malloc/tcinstall.sh

jemalloc:
	@echo "Installing jemalloc..."
	@bash ./libs/malloc/jeinstall.sh
