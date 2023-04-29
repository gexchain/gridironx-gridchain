//go:build ignore

package keeper

import (
	"testing"

	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	"github.com/gridironx/gridchain/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAccountInvariant(t *testing.T) {

	testInput := createTestInputWithBalance(t, 1, 10000)
	ctx := testInput.Ctx
	keeper := testInput.DexKeeper
	accounts := testInput.TestAddrs
	keeper.SetParams(ctx, *types.DefaultParams())

	builtInTP := GetBuiltInTokenPair()
	builtInTP.Owner = accounts[0]
	err := keeper.SaveTokenPair(ctx, builtInTP)
	require.Nil(t, err)

	// deposit xxb_fury 100 fury
	depositMsg := types.NewMsgDeposit(builtInTP.Name(),
		sdk.NewDecCoin(builtInTP.QuoteAssetSymbol, sdk.NewInt(100)), accounts[0])

	err = keeper.Deposit(ctx, builtInTP.Name(), depositMsg.Depositor, depositMsg.Amount)
	require.Nil(t, err)

	// module acount balance 100fury
	// xxb_fury deposits 100 fury. withdraw info 0 fury
	invariant := ModuleAccountInvariant(keeper, keeper.supplyKeeper)
	_, broken := invariant(ctx)
	require.False(t, broken)

	// withdraw xxb_fury 50 fury
	WithdrawMsg := types.NewMsgWithdraw(builtInTP.Name(),
		sdk.NewDecCoin(builtInTP.QuoteAssetSymbol, sdk.NewInt(50)), accounts[0])

	err = keeper.Withdraw(ctx, builtInTP.Name(), WithdrawMsg.Depositor, WithdrawMsg.Amount)
	require.Nil(t, err)

	// module acount balance 100fury
	// xxb_fury deposits 50 fury. withdraw info 50 fury
	invariant = ModuleAccountInvariant(keeper, keeper.supplyKeeper)
	_, broken = invariant(ctx)
	require.False(t, broken)

}
