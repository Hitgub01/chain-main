package app

import (
	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// App implements the common methods for a Cosmos SDK-based application
// specific blockchain.
// (Amino is still needed for Ledger at the moment)
type App interface {
	// The assigned name of the app.
	Name() string

	// The application types codec.
	// NOTE: This should be sealed before being returned.
	LegacyAmino() *codec.LegacyAmino

	// Application updates every begin block.
	BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error)

	// Application updates every end block.
	EndBlocker(ctx sdk.Context) (sdk.EndBlock, error)

	// Application update at chain (i.e app) initialization.
	InitChainer(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error)

	// Loads the app at a given height.
	LoadHeight(height int64) error

	// Exports the state of the application for a genesis file.
	ExportAppStateAndValidators(
		forZeroHeight bool, jailAllowedAddrs, modulesToExport []string,
	) (servertypes.ExportedApp, error)

	// All the registered module account addresses.
	ModuleAccountAddrs() map[string]bool
}
