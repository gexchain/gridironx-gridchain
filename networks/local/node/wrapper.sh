#!/usr/bin/env sh

##
## Input parameters
##
ID=${ID:-0}
LOG=${LOG:-gridchaind.log}

##
## Run binary with all parameters
##
export GRIDCHAINDHOME="/gridchaind/node${ID}/gridchaind"

if [ -d "$(dirname "${GRIDCHAINDHOME}"/"${LOG}")" ]; then
  gridchaind --chain-id gridchain-1 --home "${GRIDCHAINDHOME}" "$@" | tee "${GRIDCHAINDHOME}/${LOG}"
else
  gridchaind --chain-id gridchain-1 --home "${GRIDCHAINDHOME}" "$@"
fi

