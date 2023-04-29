package auth

import (
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/auth/exported"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/auth/keeper"
)

type (
	Account       = exported.Account
	ModuleAccount = exported.ModuleAccount
	ObserverI     = keeper.ObserverI
)
