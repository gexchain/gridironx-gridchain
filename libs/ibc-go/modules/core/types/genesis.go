package types

import (
	codectypes "github.com/gridironx/gridchain/libs/cosmos-sdk/codec/types"
	clienttypes "github.com/gridironx/gridchain/libs/ibc-go/modules/core/02-client/types"
	connectiontypes "github.com/gridironx/gridchain/libs/ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/gridironx/gridchain/libs/ibc-go/modules/core/04-channel/types"
)

var _ codectypes.UnpackInterfacesMessage = GenesisState{}

// DefaultGenesisState returns the ibc module's default genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		ClientGenesis:     clienttypes.DefaultGenesisState(),
		ConnectionGenesis: connectiontypes.DefaultGenesisState(),
		ChannelGenesis:    channeltypes.DefaultGenesisState(),
		Params:            DefaultParams(),
	}
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (gs GenesisState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return gs.ClientGenesis.UnpackInterfaces(unpacker)
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs *GenesisState) Validate() error {
	if err := gs.ClientGenesis.Validate(); err != nil {
		return err
	}

	if err := gs.ConnectionGenesis.Validate(); err != nil {
		return err
	}

	return gs.ChannelGenesis.Validate()
}
