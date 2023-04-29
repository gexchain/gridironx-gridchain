//go:build ignore

package keeper

import (
	"fmt"
	"testing"

	"github.com/gridironx/gridchain/libs/cosmos-sdk/codec"
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/auth"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/bank"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/mock"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/supply"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/supply/exported"
	abci "github.com/gridironx/gridchain/libs/tendermint/abci/types"
	"github.com/gridironx/gridchain/libs/tendermint/crypto/secp256k1"
	"github.com/gridironx/gridchain/x/ammswap/types"
	staking "github.com/gridironx/gridchain/x/staking/types"
	"github.com/stretchr/testify/require"

	"github.com/gridironx/gridchain/x/token"
)

type TestInput struct {
	*mock.App

	keySwap   *sdk.KVStoreKey
	keyToken  *sdk.KVStoreKey
	keyLock   *sdk.KVStoreKey
	keySupply *sdk.KVStoreKey

	bankKeeper   bank.Keeper
	swapKeeper   Keeper
	tokenKeeper  token.Keeper
	supplyKeeper supply.Keeper
}

func regCodec(cdc *codec.Codec) {
	types.RegisterCodec(cdc)
	token.RegisterCodec(cdc)
	supply.RegisterCodec(cdc)
}

func GetTestInput(t *testing.T, numGenAccs int) (mockApp *TestInput, addrKeysSlice mock.AddrKeysSlice) {
	return getTestInputWithBalance(t, numGenAccs, 100)
}

// initialize the mock application for this module
func getTestInputWithBalance(t *testing.T, numGenAccs int, balance int64) (mockApp *TestInput,
	addrKeysSlice mock.AddrKeysSlice) {
	mapp := mock.NewApp()
	regCodec(mapp.Cdc.GetCdc())

	mockApp = &TestInput{
		App:       mapp,
		keySwap:   sdk.NewKVStoreKey(types.StoreKey),
		keyToken:  sdk.NewKVStoreKey(token.StoreKey),
		keyLock:   sdk.NewKVStoreKey(token.KeyLock),
		keySupply: sdk.NewKVStoreKey(supply.StoreKey),
	}

	feeCollector := supply.NewEmptyModuleAccount(auth.FeeCollectorName)
	blacklistedAddrs := make(map[string]bool)
	blacklistedAddrs[feeCollector.String()] = true

	mockApp.bankKeeper = bank.NewBaseKeeper(mockApp.AccountKeeper,
		mockApp.ParamsKeeper.Subspace(bank.DefaultParamspace),
		blacklistedAddrs)

	maccPerms := map[string][]string{
		auth.FeeCollectorName: nil,
		token.ModuleName:      {supply.Minter, supply.Burner},
		types.ModuleName:      {supply.Minter, supply.Burner},
	}
	mockApp.supplyKeeper = supply.NewKeeper(mockApp.Cdc.GetCdc(), mockApp.keySupply, mockApp.AccountKeeper,
		mockApp.bankKeeper, maccPerms)

	mockApp.tokenKeeper = token.NewKeeper(
		mockApp.bankKeeper,
		mockApp.ParamsKeeper.Subspace(token.DefaultParamspace),
		auth.FeeCollectorName,
		mockApp.supplyKeeper,
		mockApp.keyToken,
		mockApp.keyLock,
		mockApp.Cdc.GetCdc(),
		true, mockApp.AccountKeeper)

	mockApp.swapKeeper = NewKeeper(
		mockApp.supplyKeeper,
		mockApp.tokenKeeper,
		mockApp.Cdc.GetCdc(),
		mockApp.keySwap,
		mockApp.ParamsKeeper.Subspace(types.DefaultParamspace),
	)

	mockApp.QueryRouter().AddRoute(types.QuerierRoute, NewQuerier(mockApp.swapKeeper))

	mockApp.SetInitChainer(initChainer(mockApp.App, mockApp.supplyKeeper,
		[]exported.ModuleAccountI{feeCollector}))

	decCoins, err := sdk.ParseDecCoins(fmt.Sprintf("%d%s,%d%s,%d%s,%d%s",
		balance, types.TestQuotePooledToken, balance, types.TestBasePooledToken, balance, types.TestBasePooledToken2, balance, types.TestBasePooledToken3))
	require.Nil(t, err)
	coins := decCoins

	keysSlice, genAccs := GenAccounts(numGenAccs, coins)
	addrKeysSlice = keysSlice

	// todo: checkTx in mock app
	mockApp.SetAnteHandler(nil)

	app := mockApp
	require.NoError(t, app.CompleteSetup(
		app.keySwap,
		app.keyToken,
		app.keyLock,
		app.keySupply,
	))
	mock.SetGenesis(mockApp.App, genAccs)

	for i := 0; i < numGenAccs; i++ {
		mock.CheckBalance(t, app.App, keysSlice[i].Address, coins)
		mockApp.TotalCoinsSupply = mockApp.TotalCoinsSupply.Add2(coins)
	}

	return mockApp, addrKeysSlice
}

func initChainer(mapp *mock.App, supplyKeeper staking.SupplyKeeper,
	blacklistedAddrs []exported.ModuleAccountI) sdk.InitChainer {
	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
		mapp.InitChainer(ctx, req)
		// set module accounts
		for _, macc := range blacklistedAddrs {
			supplyKeeper.SetModuleAccount(ctx, macc)
		}
		return abci.ResponseInitChain{}
	}
}

func GenAccounts(numAccs int, genCoins sdk.Coins) (addrKeysSlice mock.AddrKeysSlice,
	genAccs []auth.Account) {
	for i := 0; i < numAccs; i++ {
		privKey := secp256k1.GenPrivKey()
		pubKey := privKey.PubKey()
		addr := sdk.AccAddress(pubKey.Address())

		addrKeys := mock.NewAddrKeys(addr, pubKey, privKey)
		account := &auth.BaseAccount{
			Address: addr,
			Coins:   genCoins,
		}
		genAccs = append(genAccs, account)
		addrKeysSlice = append(addrKeysSlice, addrKeys)
	}
	return
}
