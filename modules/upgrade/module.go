package upgrade

import (
	"github.com/villagelabsco/bdjuno/v3/database"

	"github.com/villagelabsco/juno/v4/modules"
)

var (
	_ modules.Module      = &Module{}
	_ modules.BlockModule = &Module{}
)

// Module represents the x/upgrade module
type Module struct {
	db            *database.Db
	stakingModule StakingModule
}

// NewModule builds a new Module instance
func NewModule(db *database.Db, stakingModule StakingModule) *Module {
	return &Module{
		stakingModule: stakingModule,
		db:            db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "upgrade"
}
