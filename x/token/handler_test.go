package token_test

import (
	"testing"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	gridchain "github.com/gridironx/gridchain/app"
	app "github.com/gridironx/gridchain/app/types"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/codec"
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/auth"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/mock"
	abci "github.com/gridironx/gridchain/libs/tendermint/abci/types"
	"github.com/gridironx/gridchain/libs/tendermint/crypto/secp256k1"
	"github.com/gridironx/gridchain/libs/tendermint/libs/log"
	dbm "github.com/gridironx/gridchain/libs/tm-db"
	"github.com/gridironx/gridchain/x/common"
	"github.com/gridironx/gridchain/x/common/version"
	"github.com/gridironx/gridchain/x/token"
	"github.com/gridironx/gridchain/x/token/types"
	"github.com/stretchr/testify/require"
)

func TestHandlerBlockedContractAddrSend(t *testing.T) {
	gridironxapp := initApp(true)
	ctx := gridironxapp.BaseApp.NewContext(true, abci.Header{Height: 1})
	gAcc := CreateEthAccounts(3, sdk.SysCoins{
		sdk.NewDecCoinFromDec(common.NativeToken, sdk.NewDec(10000)),
	})
	gridironxapp.AccountKeeper.SetAccount(ctx, gAcc[0])
	gridironxapp.AccountKeeper.SetAccount(ctx, gAcc[1])
	gAcc[2].CodeHash = []byte("contract code hash")
	gridironxapp.AccountKeeper.SetAccount(ctx, gAcc[2])

	// multi send
	multiSendStr := `[{"to":"` + gAcc[1].Address.String() + `","amount":" 10` + common.NativeToken + `"}]`
	transfers, err := types.StrToTransfers(multiSendStr)
	require.Nil(t, err)
	multiSendStr2 := `[{"to":"` + gAcc[2].Address.String() + `","amount":" 10` + common.NativeToken + `"}]`
	transfers2, err := types.StrToTransfers(multiSendStr2)
	require.Nil(t, err)

	successfulSendMsg := types.NewMsgTokenSend(gAcc[0].Address, gAcc[1].Address, sdk.SysCoins{sdk.NewDecCoinFromDec(common.NativeToken, sdk.NewDec(1))})
	sendToContractMsg := types.NewMsgTokenSend(gAcc[0].Address, gAcc[2].Address, sdk.SysCoins{sdk.NewDecCoinFromDec(common.NativeToken, sdk.NewDec(1))})
	successfulMultiSendMsg := types.NewMsgMultiSend(gAcc[0].Address, transfers)
	multiSendToContractMsg := types.NewMsgMultiSend(gAcc[0].Address, transfers2)
	handler := token.NewTokenHandler(gridironxapp.TokenKeeper, version.CurrentProtocolVersion)
	gridironxapp.BankKeeper.SetSendEnabled(ctx, true)
	TestSets := []struct {
		description string
		balance     string
		msg         sdk.Msg
		account     app.EthAccount
	}{
		// 0.01fury as fixed fee in each stdTx
		{"success to send", "9999.000000000000000000fury", successfulSendMsg, gAcc[0]},
		{"success to multi-send", "9989.000000000000000000fury", successfulMultiSendMsg, gAcc[0]},
		{"success to send", "9988.000000000000000000fury", successfulSendMsg, gAcc[0]},
		{"success to multi-send", "9978.000000000000000000fury", successfulMultiSendMsg, gAcc[0]},
		//{"fail to send to contract", "9978.000000000000000000fury", failedSendMsg, gAcc[0]},
		//{"fail to multi-send to contract", "9978.000000000000000000fury", failedMultiSendMsg, gAcc[0]},
		{"fail to send to contract", "9978.000000000000000000fury", sendToContractMsg, gAcc[0]},
		{"fail to multi-send to contract", "9978.000000000000000000fury", multiSendToContractMsg, gAcc[0]},
	}
	for i, tt := range TestSets {
		t.Run(tt.description, func(t *testing.T) {
			handler(ctx, TestSets[i].msg)
			acc := gridironxapp.AccountKeeper.GetAccount(ctx, tt.account.Address)
			acc.GetCoins().String()
			require.Equal(t, acc.GetCoins().String(), tt.balance)
		})
	}
}

// Setup initializes a new GRIDIronxChainApp. A Nop logger is set in GRIDIronxChainApp.
func initApp(isCheckTx bool) *gridchain.GRIDIronxChainApp {
	db := dbm.NewMemDB()
	app := gridchain.NewGRIDIronxChainApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, 0)

	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		genesisState := gridchain.NewDefaultGenesisState()
		stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:    []abci.ValidatorUpdate{},
				AppStateBytes: stateBytes,
			},
		)
		app.EndBlock(abci.RequestEndBlock{})
		app.Commit(abci.RequestCommit{})
	}

	return app
}

func CreateEthAccounts(numAccs int, genCoins sdk.SysCoins) (genAccs []app.EthAccount) {
	for i := 0; i < numAccs; i++ {
		privKey := secp256k1.GenPrivKey()
		pubKey := privKey.PubKey()
		addr := sdk.AccAddress(pubKey.Address())

		ak := mock.NewAddrKeys(addr, pubKey, privKey)
		testAccount := app.EthAccount{
			BaseAccount: &auth.BaseAccount{
				Address: ak.Address,
				Coins:   genCoins,
			},
			CodeHash: ethcrypto.Keccak256(nil),
		}
		genAccs = append(genAccs, testAccount)
	}
	return
}
