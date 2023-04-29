#!/usr/bin/env sh

##
## Input parameters
##
ID=${ID:-0}
LOG=${LOG:-gridchaind.log}

##
## Run binary with all parameters
##
export EXCHAINDHOME="/gridchaind/node${ID}/gridchaind"

if [ -d "$(dirname "${EXCHAINDHOME}"/"${LOG}")" ]; then
  gridchaind --chain-id gridchain-1 --home "${EXCHAINDHOME}" "$@" | tee "${EXCHAINDHOME}/${LOG}"
else
  gridchaind --chain-id gridchain-1 --home "${EXCHAINDHOME}" "$@"
fi

