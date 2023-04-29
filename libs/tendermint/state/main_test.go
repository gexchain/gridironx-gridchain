package state_test

import (
	"os"
	"testing"

	"github.com/gridironx/gridchain/libs/tendermint/types"
)

func TestMain(m *testing.M) {
	types.RegisterMockEvidencesGlobal()
	os.Exit(m.Run())
}
