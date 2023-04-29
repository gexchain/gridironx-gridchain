res=$(gridchaincli tx wasm store ./wasm/escrow/artifacts/cw_escrow-aarch64.wasm --fees 0.01fury --from captain --gas=2000000 -b block -y)
echo "store code..."
echo $res
code_id=$(echo "$res" | jq '.logs[0].events[1].attributes[0].value' | sed 's/\"//g')
res=$(gridchaincli tx wasm instantiate "$code_id" '{"arbiter":"0xbbE4733d85bc2b90682147779DA49caB38C0aA1F","end_height":100000,"recipient":"0x2Bd4AF0C1D0c2930fEE852D07bB9dE87D8C07044"}' --label test1 --admin did:fury:ex1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v --fees 0.001fury --from captain -b block -y)
contractAddr=$(echo "$res" | jq '.logs[0].events[0].attributes[0].value' | sed 's/\"//g')
echo "instantiate contract..."
echo $res
#gridchaincli tx send did:fury:ex1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v $contractAddr 999fury --fees 0.01fury -y -b block
gridchaincli tx wasm execute "$contractAddr" '{"approve":{"quantity":[{"amount":"1","denom":"fury"}]}}' --amount 888fury --fees 0.001fury --from captain -b block -y
