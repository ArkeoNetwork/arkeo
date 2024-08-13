package upgrades

import (
	"cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/arkeonetwork/arkeo/app/keepers"
	"github.com/cosmos/cosmos-sdk/types/module"
)

type Upgrade struct {
	UpgradeName          string
	CreateUpgradeHandler func(*module.Manager, module.Configurator, keepers.ArkeoKeepers) upgradetypes.UpgradeHandler
	StoreUpgrades        types.StoreUpgrades
}