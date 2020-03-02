package version

import (
	"github.com/filecoin-project/specs-actors/actors/abi"
)

// USER is the user network
const USER = "alpha2"

// DEVNET4 is the network name of devnet
const DEVNET4 = "devnet4"

// LOCALNET is the network name of localnet
const LOCALNET = "localnet"

// TESTNET3 is the network name of the third testnet
const TESTNET3 = "localnet-270a8688-1b23-4508-b675-444cb1e6f05d"

// TEST is the network name for internal tests
const TEST = "gfctest"

// Protocol0 is the first protocol version
const Protocol0 = 0

// Protocol1 is the weight upgrade
const Protocol1 = 1

// ConfigureProtocolVersions configures all protocol upgrades for all known networks.
// TODO: support arbitrary network names at "latest" protocol version so that only coordinated
// network upgrades need to be represented here. See #3491.
func ConfigureProtocolVersions(network string) (*ProtocolVersionTable, error) {
	return NewProtocolVersionTableBuilder(network).
		Add(USER, Protocol0, abi.ChainEpoch(0)).
		Add(USER, Protocol1, abi.ChainEpoch(43000)).
		Add(DEVNET4, Protocol0, abi.ChainEpoch(0)).
		Add(DEVNET4, Protocol1, abi.ChainEpoch(300)).
		Add(LOCALNET, Protocol1, abi.ChainEpoch(0)).
		Add(TESTNET3, Protocol1, abi.ChainEpoch(0)).
		Add(TEST, Protocol1, abi.ChainEpoch(0)).
		Build()
}
