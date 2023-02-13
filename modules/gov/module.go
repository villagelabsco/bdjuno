package gov

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/villagelabsco/bdjuno/v3/database"

	govsource "github.com/villagelabsco/bdjuno/v3/modules/gov/source"

	"github.com/villagelabsco/juno/v4/modules"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.BlockModule   = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represent x/gov module
type Module struct {
	cdc            codec.Codec
	db             *database.Db
	source         govsource.Source
	authModule     AuthModule
	distrModule    DistrModule
	mintModule     MintModule
	slashingModule SlashingModule
	stakingModule  StakingModule
	identityModule IdentityModule
}

// NewModule returns a new Module instance
func NewModule(
	source govsource.Source,
	authModule AuthModule,
	distrModule DistrModule,
	mintModule MintModule,
	slashingModule SlashingModule,
	stakingModule StakingModule,
	identityModule IdentityModule,
	cdc codec.Codec,
	db *database.Db,
) *Module {
	return &Module{
		cdc:            cdc,
		db:             db,
		source:         source,
		authModule:     authModule,
		distrModule:    distrModule,
		mintModule:     mintModule,
		slashingModule: slashingModule,
		stakingModule:  stakingModule,
		identityModule: identityModule,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "gov"
}
