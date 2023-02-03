package reputation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/villagelabsco/bdjuno/v3/database"
	reputationsource "github.com/villagelabsco/bdjuno/v3/modules/reputation/source"
	"github.com/villagelabsco/juno/v4/modules"
	reputationtypes "github.com/villagelabsco/villaged/x/reputation/types"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represent x/feegrant module
type Module struct {
	cdc codec.Codec
	db  *database.Db
	s   reputationsource.Source
}

// NewModule returns a new Module instance
func NewModule(cdc codec.Codec, db *database.Db, s reputationsource.Source) *Module {
	return &Module{
		cdc: cdc,
		db:  db,
		s:   s,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return reputationtypes.ModuleName
}
