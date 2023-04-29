package types

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
	"testing"
	"time"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContractInfoValidateBasic(t *testing.T) {
	specs := map[string]struct {
		srcMutator func(*ContractInfo)
		expError   bool
	}{
		"all good": {srcMutator: func(_ *ContractInfo) {}},
		"code id empty": {
			srcMutator: func(c *ContractInfo) { c.CodeID = 0 },
			expError:   true,
		},
		"creator empty": {
			srcMutator: func(c *ContractInfo) { c.Creator = "" },
			expError:   true,
		},
		"creator not an address": {
			srcMutator: func(c *ContractInfo) { c.Creator = "invalid address" },
			expError:   true,
		},
		"admin empty": {
			srcMutator: func(c *ContractInfo) { c.Admin = "" },
			expError:   false,
		},
		"admin not an address": {
			srcMutator: func(c *ContractInfo) { c.Admin = "invalid address" },
			expError:   true,
		},
		"label empty": {
			srcMutator: func(c *ContractInfo) { c.Label = "" },
			expError:   true,
		},
		"label exceeds limit": {
			srcMutator: func(c *ContractInfo) { c.Label = strings.Repeat("a", MaxLabelSize+1) },
			expError:   true,
		},
		//"invalid extension": {
		//	srcMutator: func(c *ContractInfo) {
		//		// any protobuf type with ValidateBasic method
		//		any, err := codectypes.NewAnyWithValue(&govtypes.TextProposal{})
		//		require.NoError(t, err)
		//		c.Extension = any
		//	},
		//	expError: true,
		//},
		//"not validatable extension": {
		//	srcMutator: func(c *ContractInfo) {
		//		// any protobuf type with ValidateBasic method
		//		any, err := codectypes.NewAnyWithValue(&govtypes.Proposal{})
		//		require.NoError(t, err)
		//		c.Extension = any
		//	},
		//},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			state := ContractInfoFixture(spec.srcMutator)
			got := state.ValidateBasic()
			if spec.expError {
				require.Error(t, got)
				return
			}
			require.NoError(t, got)
		})
	}
}

func TestCodeInfoValidateBasic(t *testing.T) {
	specs := map[string]struct {
		srcMutator func(*CodeInfo)
		expError   bool
	}{
		"all good": {srcMutator: func(_ *CodeInfo) {}},
		"code hash empty": {
			srcMutator: func(c *CodeInfo) { c.CodeHash = []byte{} },
			expError:   true,
		},
		"code hash nil": {
			srcMutator: func(c *CodeInfo) { c.CodeHash = nil },
			expError:   true,
		},
		"creator empty": {
			srcMutator: func(c *CodeInfo) { c.Creator = "" },
			expError:   true,
		},
		"creator not an address": {
			srcMutator: func(c *CodeInfo) { c.Creator = "invalid address" },
			expError:   true,
		},
		"Instantiate config invalid": {
			srcMutator: func(c *CodeInfo) { c.InstantiateConfig = AccessConfig{} },
			expError:   true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			state := CodeInfoFixture(spec.srcMutator)
			got := state.ValidateBasic()
			if spec.expError {
				require.Error(t, got)
				return
			}
			require.NoError(t, got)
		})
	}
}

//func TestContractInfoSetExtension(t *testing.T) {
//	anyTime := time.Now().UTC()
//	aNestedProtobufExt := func() ContractInfoExtension {
//		// using gov proposal here as a random protobuf types as it contains an Any type inside for nested unpacking
//		myExtension, err := govtypes.NewProposal(&govtypes.TextProposal{Title: "bar"}, 1, anyTime, anyTime)
//		require.NoError(t, err)
//		myExtension.TotalDeposit = nil
//		return &myExtension
//	}
//
//	specs := map[string]struct {
//		src    ContractInfoExtension
//		expErr bool
//		expNil bool
//	}{
//		"all good with any proto extension": {
//			src: aNestedProtobufExt(),
//		},
//		"nil allowed": {
//			src:    nil,
//			expNil: true,
//		},
//		"validated and accepted": {
//			src: &govtypes.TextProposal{Title: "bar", Description: "set"},
//		},
//		"validated and rejected": {
//			src:    &govtypes.TextProposal{Title: "bar"},
//			expErr: true,
//		},
//	}
//	for name, spec := range specs {
//		t.Run(name, func(t *testing.T) {
//			var c ContractInfo
//			gotErr := c.SetExtension(spec.src)
//			if spec.expErr {
//				require.Error(t, gotErr)
//				return
//			}
//			require.NoError(t, gotErr)
//			if spec.expNil {
//				return
//			}
//			require.NotNil(t, c.Extension)
//			assert.NotNil(t, c.Extension.GetCachedValue())
//		})
//	}
//}
//
//func TestContractInfoMarshalUnmarshal(t *testing.T) {
//	var myAddr sdk.WasmAddress = rand.Bytes(ContractAddrLen)
//	var myOtherAddr sdk.WasmAddress = rand.Bytes(ContractAddrLen)
//	anyPos := AbsoluteTxPosition{BlockHeight: 1, TxIndex: 2}
//
//	anyTime := time.Now().UTC()
//	// using gov proposal here as a random protobuf types as it contains an Any type inside for nested unpacking
//	myExtension, err := govtypes.NewProposal(&govtypes.TextProposal{Title: "bar"}, 1, anyTime, anyTime)
//	require.NoError(t, err)
//	myExtension.TotalDeposit = nil
//
//	src := NewContractInfo(1, myAddr, myOtherAddr, "bar", &anyPos)
//	err = src.SetExtension(&myExtension)
//	require.NoError(t, err)
//
//	interfaceRegistry := types.NewInterfaceRegistry()
//	marshaler := codec.NewProtoCodec(interfaceRegistry)
//	RegisterInterfaces(interfaceRegistry)
//	// register proposal as extension type
//	interfaceRegistry.RegisterImplementations(
//		(*ContractInfoExtension)(nil),
//		&govtypes.Proposal{},
//	)
//	// register gov types for nested Anys
//	govtypes.RegisterInterfaces(interfaceRegistry)
//
//	// when encode
//	bz, err := marshaler.Marshal(&src)
//	require.NoError(t, err)
//	// and decode
//	var dest ContractInfo
//	err = marshaler.Unmarshal(bz, &dest)
//	// then
//	require.NoError(t, err)
//	assert.Equal(t, src, dest)
//	// and sanity check nested any
//	var destExt govtypes.Proposal
//	require.NoError(t, dest.ReadExtension(&destExt))
//	assert.Equal(t, destExt.GetTitle(), "bar")
//}
//
//func TestContractInfoReadExtension(t *testing.T) {
//	anyTime := time.Now().UTC()
//	myExtension, err := govtypes.NewProposal(&govtypes.TextProposal{Title: "foo"}, 1, anyTime, anyTime)
//	require.NoError(t, err)
//	type TestExtensionAsStruct struct {
//		ContractInfoExtension
//	}
//
//	specs := map[string]struct {
//		setup  func(*ContractInfo)
//		param  func() ContractInfoExtension
//		expVal ContractInfoExtension
//		expErr bool
//	}{
//		"all good": {
//			setup: func(i *ContractInfo) {
//				i.SetExtension(&myExtension)
//			},
//			param: func() ContractInfoExtension {
//				return &govtypes.Proposal{}
//			},
//			expVal: &myExtension,
//		},
//		"no extension set": {
//			setup: func(i *ContractInfo) {
//			},
//			param: func() ContractInfoExtension {
//				return &govtypes.Proposal{}
//			},
//			expVal: &govtypes.Proposal{},
//		},
//		"nil argument value": {
//			setup: func(i *ContractInfo) {
//				i.SetExtension(&myExtension)
//			},
//			param: func() ContractInfoExtension {
//				return nil
//			},
//			expErr: true,
//		},
//		"non matching types": {
//			setup: func(i *ContractInfo) {
//				i.SetExtension(&myExtension)
//			},
//			param: func() ContractInfoExtension {
//				return &govtypes.TextProposal{}
//			},
//			expErr: true,
//		},
//		"not a pointer argument": {
//			setup: func(i *ContractInfo) {
//			},
//			param: func() ContractInfoExtension {
//				return TestExtensionAsStruct{}
//			},
//			expErr: true,
//		},
//	}
//	for name, spec := range specs {
//		t.Run(name, func(t *testing.T) {
//			var c ContractInfo
//			spec.setup(&c)
//			// when
//
//			gotValue := spec.param()
//			gotErr := c.ReadExtension(gotValue)
//
//			// then
//			if spec.expErr {
//				require.Error(t, gotErr)
//				return
//			}
//			require.NoError(t, gotErr)
//			assert.Equal(t, spec.expVal, gotValue)
//		})
//	}
//}

func TestNewEnv(t *testing.T) {
	myTime := time.Unix(0, 1619700924259075000)
	ctx := (&sdk.Context{}).SetChainID("testing").SetContext(context.Background())
	t.Logf("++ unix: %d", myTime.UnixNano())
	var myContractAddr sdk.WasmAddress = randBytes(SDKAddrLen)
	specs := map[string]struct {
		srcCtx sdk.Context
		exp    wasmvmtypes.Env
	}{
		"all good with tx counter": {
			srcCtx: WithTXCounter((*ctx).WithBlockHeight(1).WithBlockTime(myTime), 0),
			exp: wasmvmtypes.Env{
				Block: wasmvmtypes.BlockInfo{
					Height:  1,
					Time:    1619700924259075000,
					ChainID: "testing",
				},
				Contract: wasmvmtypes.ContractInfo{
					Address: myContractAddr.String(),
				},
				Transaction: &wasmvmtypes.TransactionInfo{Index: 0},
			},
		},
		//"without tx counter": {
		//	srcCtx: (*ctx).WithBlockHeight(1).WithBlockTime(myTime),
		//	exp: wasmvmtypes.Env{
		//		Block: wasmvmtypes.BlockInfo{
		//			Height:  1,
		//			Time:    1619700924259075000,
		//			ChainID: "testing",
		//		},
		//		Contract: wasmvmtypes.ContractInfo{
		//			Address: myContractAddr.String(),
		//		},
		//	},
		//},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, spec.exp, NewEnv(spec.srcCtx, myContractAddr))
		})
	}
}

func TestVerifyAddressLen(t *testing.T) {
	specs := map[string]struct {
		src    []byte
		expErr bool
	}{
		"valid contract address": {
			src: bytes.Repeat([]byte{1}, 20),
		},
		"valid legacy address": {
			src: bytes.Repeat([]byte{1}, 20),
		},
		"address too short for legacy": {
			src:    bytes.Repeat([]byte{1}, 19),
			expErr: true,
		},
		"address too short for contract": {
			src:    bytes.Repeat([]byte{1}, 19),
			expErr: true,
		},
		"address too long for legacy": {
			src:    bytes.Repeat([]byte{1}, 21),
			expErr: true,
		},
		"address too long for contract": {
			src:    bytes.Repeat([]byte{1}, 21),
			expErr: true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			gotErr := VerifyAddressLen()(spec.src)
			if spec.expErr {
				require.Error(t, gotErr)
				return
			}
			require.NoError(t, gotErr)
		})
	}
}

func TestAccesConfigSubset(t *testing.T) {
	specs := map[string]struct {
		check    AccessConfig
		superSet AccessConfig
		isSubSet bool
	}{
		"nobody <= nobody": {
			superSet: AccessConfig{Permission: AccessTypeNobody},
			check:    AccessConfig{Permission: AccessTypeNobody},
			isSubSet: true,
		},
		"only > nobody": {
			superSet: AccessConfig{Permission: AccessTypeNobody},
			check:    AccessConfig{Permission: AccessTypeOnlyAddress, Address: "foobar"},
			isSubSet: false,
		},
		"everybody > nobody": {
			superSet: AccessConfig{Permission: AccessTypeNobody},
			check:    AccessConfig{Permission: AccessTypeEverybody},
			isSubSet: false,
		},
		"unspecified > nobody": {
			superSet: AccessConfig{Permission: AccessTypeNobody},
			check:    AccessConfig{Permission: AccessTypeUnspecified},
			isSubSet: false,
		},
		"nobody <= everybody": {
			superSet: AccessConfig{Permission: AccessTypeEverybody},
			check:    AccessConfig{Permission: AccessTypeNobody},
			isSubSet: true,
		},
		"only <= everybody": {
			superSet: AccessConfig{Permission: AccessTypeEverybody},
			check:    AccessConfig{Permission: AccessTypeOnlyAddress, Address: "foobar"},
			isSubSet: true,
		},
		"everybody <= everybody": {
			superSet: AccessConfig{Permission: AccessTypeEverybody},
			check:    AccessConfig{Permission: AccessTypeEverybody},
			isSubSet: true,
		},
		"unspecified > everybody": {
			superSet: AccessConfig{Permission: AccessTypeEverybody},
			check:    AccessConfig{Permission: AccessTypeUnspecified},
			isSubSet: false,
		},
		"nobody <= only": {
			superSet: AccessConfig{Permission: AccessTypeOnlyAddress, Address: "owner"},
			check:    AccessConfig{Permission: AccessTypeNobody},
			isSubSet: true,
		},
		"only <= only(same)": {
			superSet: AccessConfig{Permission: AccessTypeOnlyAddress, Address: "owner"},
			check:    AccessConfig{Permission: AccessTypeOnlyAddress, Address: "owner"},
			isSubSet: true,
		},
		"only > only(other)": {
			superSet: AccessConfig{Permission: AccessTypeOnlyAddress, Address: "owner"},
			check:    AccessConfig{Permission: AccessTypeOnlyAddress, Address: "other"},
			isSubSet: false,
		},
		"everybody > only": {
			superSet: AccessConfig{Permission: AccessTypeOnlyAddress, Address: "owner"},
			check:    AccessConfig{Permission: AccessTypeEverybody},
			isSubSet: false,
		},
		"nobody > unspecified": {
			superSet: AccessConfig{Permission: AccessTypeUnspecified},
			check:    AccessConfig{Permission: AccessTypeNobody},
			isSubSet: false,
		},
	}

	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			subset := spec.check.IsSubset(spec.superSet)
			require.Equal(t, spec.isSubSet, subset)
		})
	}
}

func TestConvertAccessConfig(t *testing.T) {
	did:fury:ex1, eth1 := generateAddress(1)
	ex2, eth2 := generateAddress(1)

	testcase := []struct {
		name   string
		config AccessConfig
		expect AccessConfig
		isErr  bool
	}{
		{
			name:   "nobody",
			config: AccessConfig{Permission: AccessTypeNobody},
			expect: AccessConfig{Permission: AccessTypeNobody},
			isErr:  false,
		},
		{
			name:   "everybody",
			config: AccessConfig{Permission: AccessTypeEverybody},
			expect: AccessConfig{Permission: AccessTypeEverybody},
			isErr:  false,
		},
		{
			name:   "only 1 ethaddr",
			config: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{eth1.String()}, ",")},
			expect: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{eth1.String()}, ",")},
			isErr:  false,
		},
		{
			name:   "only 2 ethaddr",
			config: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{eth1.String(), eth2.String()}, ",")},
			expect: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{eth1.String(), eth2.String()}, ",")},
			isErr:  false,
		},
		{
			name:   "only 1 exaddr",
			config: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{did:fury:ex1.String()}, ",")},
			expect: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{eth1.String()}, ",")},
			isErr:  false,
		},
		{
			name:   "only 2 exaddr",
			config: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{did:fury:ex1.String(), ex2.String()}, ",")},
			expect: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{eth1.String(), eth2.String()}, ",")},
			isErr:  false,
		},
		{
			name:   "only 1 exaddr 1 ethaddr",
			config: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{eth1.String(), ex2.String()}, ",")},
			expect: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{eth1.String(), eth2.String()}, ",")},
			isErr:  false,
		},
		{
			name:   "only 1 exaddr 1 erraddr",
			config: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{eth1.String(), "erraddr"}, ",")},
			expect: AccessConfig{},
			isErr:  true,
		},
		{
			name:   "only  1 erraddr",
			config: AccessConfig{Permission: AccessTypeOnlyAddress, Address: strings.Join([]string{"erraddr"}, ",")},
			expect: AccessConfig{},
			isErr:  true,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ConvertAccessConfig(tc.config)
			if tc.isErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expect, result)
			}

		})
	}
}

func generateAddress(seed int64) (sdk.AccAddress, common.Address) {
	buff := crypto.Keccak256Hash(big.NewInt(seed).Bytes()).Bytes()[:20]
	return sdk.AccAddress(buff), common.BytesToAddress(buff)
}
