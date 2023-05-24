package reputation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v4/modules"
	"github.com/villagelabsco/bdjuno/v4/database"
	reputationsource "github.com/villagelabsco/bdjuno/v4/modules/reputation/source"
	reputationtypes "github.com/villagelabsco/village/x/reputation/types"
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
