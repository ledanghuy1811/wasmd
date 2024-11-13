package v050

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/CosmWasm/wasmd/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name
const UpgradeName = "v0.50.1"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        storetypes.StoreUpgrades{
		// Added: []string{
		// 	circuittypes.ModuleName,
		// 	consensustypes.StoreKey,
		// 	crisistypes.StoreKey,
		// 	erc20types.StoreKey,
		// 	group.StoreKey,
		// },
		// Deleted: []string{"utilevm", "evmutil", "intertx"},
	},
}

func CreateUpgradeHandler(
	mm upgrades.ModuleManager,
	configurator module.Configurator,
	ak *upgrades.AppKeepers,
) upgradetypes.UpgradeHandler {
	// sdk 47 to sdk 50
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
