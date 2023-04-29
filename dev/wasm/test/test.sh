#!/bin/bash
set -o errexit -o nounset -o pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

gridchaincli keys add fred
echo "0-----------------------"
gridchaincli tx send captain $(gridchaincli keys show fred -a) 1000fury --fees 0.001fury -y -b block

echo "1-----------------------"
echo "## Add new CosmWasm contract"
RESP=$(gridchaincli tx wasm store "$DIR/../../../x/wasm/keeper/testdata/hackatom.wasm" \
  --from captain --fees 0.001fury --gas 1500000 -y --node=http://localhost:26657 -b block -o json)

CODE_ID=$(echo "$RESP" | jq -r '.logs[0].events[1].attributes[-1].value')
echo "* Code id: $CODE_ID"
echo "* Download code"
TMPDIR=$(mktemp -t gridchaincliXXXXXX)
gridchaincli q wasm code "$CODE_ID" "$TMPDIR"
rm -f "$TMPDIR"
echo "-----------------------"
echo "## List code"
gridchaincli query wasm list-code --node=http://localhost:26657 -o json | jq

echo "2-----------------------"
echo "## Create new contract instance"
INIT="{\"verifier\":\"$(gridchaincli keys show captain | jq -r '.eth_address')\", \"beneficiary\":\"$(gridchaincli keys show fred | jq -r '.eth_address')\"}"
gridchaincli tx wasm instantiate "$CODE_ID" "$INIT" --admin="$(gridchaincli keys show captain -a)" \
  --from captain  --fees 0.001fury --amount="100fury" --label "local0.1.0" \
  --gas 1000000 -y -b block -o json | jq

CONTRACT=$(gridchaincli query wasm list-contract-by-code "$CODE_ID" -o json | jq -r '.contracts[-1]')
echo "* Contract address: $CONTRACT"
echo "### Query all"
RESP=$(gridchaincli query wasm contract-state all "$CONTRACT" -o json)
echo "$RESP" | jq
echo "### Query smart"
gridchaincli query wasm contract-state smart "$CONTRACT" '{"verifier":{}}' -o json | jq
echo "### Query raw"
KEY=$(echo "$RESP" | jq -r ".models[0].key")
gridchaincli query wasm contract-state raw "$CONTRACT" "$KEY" -o json | jq

echo "3-----------------------"
echo "## Execute contract $CONTRACT"
MSG='{"release":{}}'
gridchaincli tx wasm execute "$CONTRACT" "$MSG" \
  --from captain \
  --gas 1000000 --fees 0.001fury -y  -b block -o json | jq

echo "4-----------------------"
echo "## Set new admin"
echo "### Query old admin: $(gridchaincli q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"
echo "### Update contract"
gridchaincli tx wasm set-contract-admin "$CONTRACT" "$(gridchaincli keys show fred -a)" \
  --from captain --fees 0.001fury -y -b block -o json | jq
echo "### Query new admin: $(gridchaincli q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"

echo "5-----------------------"
echo "## Migrate contract"
echo "### Upload new code"
RESP=$(gridchaincli tx wasm store "$DIR/../../../x/wasm/keeper/testdata/burner.wasm" \
  --from captain --fees 0.001fury --gas 1000000 -y --node=http://localhost:26657 -b block -o json)

BURNER_CODE_ID=$(echo "$RESP" | jq -r '.logs[0].events[1].attributes[-1].value')
echo "### Migrate to code id: $BURNER_CODE_ID"

DEST_ACCOUNT=$(gridchaincli keys show fred | jq -r '.eth_address')
gridchaincli tx wasm migrate "$CONTRACT" "$BURNER_CODE_ID" "{\"payout\": \"$DEST_ACCOUNT\"}" --from fred  --fees 0.001fury \
 -b block -y -o json | jq

echo "### Query destination account: $BURNER_CODE_ID"
gridchaincli q account "$DEST_ACCOUNT" -o json | jq
echo "### Query contract meta data: $CONTRACT"
gridchaincli q wasm contract "$CONTRACT" -o json | jq

echo "### Query contract meta history: $CONTRACT"
gridchaincli q wasm contract-history "$CONTRACT" -o json | jq

echo "6-----------------------"
echo "## Clear contract admin"
echo "### Query old admin: $(gridchaincli q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"
echo "### Update contract"
gridchaincli tx wasm clear-contract-admin "$CONTRACT" --fees 0.001fury \
  --from fred -y -b block -o json | jq
echo "### Query new admin: $(gridchaincli q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"
