package identity

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	identitytypes "github.com/villagelabsco/village/x/identity/types"
)

// GetGenesisVestingAccounts parses the given appState and returns the genesis vesting accounts
func GetNetworksFromGenesis(appState map[string]json.RawMessage, cdc codec.Codec) ([]identitytypes.Network, error) {

	var authState identitytypes.GenesisState
	if err := cdc.UnmarshalJSON(appState["identity"], &authState); err != nil {
		return nil, err
	}

	networks := authState.NetworkList
	return networks, nil
}

// GetGenesisVestingAccounts parses the given appState and returns the genesis vesting accounts
func GetProvidersFromGenesis(appState map[string]json.RawMessage, cdc codec.Codec) ([]identitytypes.IdentityProvider, error) {

	var authState identitytypes.GenesisState
	if err := cdc.UnmarshalJSON(appState["identity"], &authState); err != nil {
		return nil, err
	}

	networks := authState.IdentityProviderList
	return networks, nil

}
