package modules

import (
	"github.com/villagelabsco/bdjuno/v3/modules/actions"
	"github.com/villagelabsco/bdjuno/v3/modules/economics"
	"github.com/villagelabsco/bdjuno/v3/modules/identity"
	"github.com/villagelabsco/bdjuno/v3/modules/marketplace"
	"github.com/villagelabsco/bdjuno/v3/modules/products"
	"github.com/villagelabsco/bdjuno/v3/modules/rbac"
	"github.com/villagelabsco/bdjuno/v3/modules/reputation"
	"github.com/villagelabsco/bdjuno/v3/modules/token"
	"github.com/villagelabsco/bdjuno/v3/modules/types"
	"github.com/villagelabsco/juno/v4/modules/pruning"
	"github.com/villagelabsco/juno/v4/modules/telemetry"

	"github.com/villagelabsco/bdjuno/v3/modules/slashing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	jmodules "github.com/villagelabsco/juno/v4/modules"
	"github.com/villagelabsco/juno/v4/modules/messages"
	"github.com/villagelabsco/juno/v4/modules/registrar"

	"github.com/villagelabsco/bdjuno/v3/utils"

	"github.com/villagelabsco/bdjuno/v3/database"
	"github.com/villagelabsco/bdjuno/v3/modules/auth"
	"github.com/villagelabsco/bdjuno/v3/modules/bank"
	"github.com/villagelabsco/bdjuno/v3/modules/consensus"
	"github.com/villagelabsco/bdjuno/v3/modules/distribution"
	"github.com/villagelabsco/bdjuno/v3/modules/feegrant"

	dailyrefetch "github.com/villagelabsco/bdjuno/v3/modules/daily_refetch"
	"github.com/villagelabsco/bdjuno/v3/modules/gov"
	"github.com/villagelabsco/bdjuno/v3/modules/mint"
	"github.com/villagelabsco/bdjuno/v3/modules/modules"
	"github.com/villagelabsco/bdjuno/v3/modules/pricefeed"
	"github.com/villagelabsco/bdjuno/v3/modules/staking"
	"github.com/villagelabsco/bdjuno/v3/modules/upgrade"
)

// UniqueAddressesParser returns a wrapper around the given parser that removes all duplicated addresses
func UniqueAddressesParser(parser messages.MessageAddressesParser) messages.MessageAddressesParser {
	return func(cdc codec.Codec, msg sdk.Msg) ([]string, error) {
		addresses, err := parser(cdc, msg)
		if err != nil {
			return nil, err
		}

		return utils.RemoveDuplicateValues(addresses), nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ registrar.Registrar = &Registrar{}
)

// Registrar represents the modules.Registrar that allows to register all modules that are supported by BigDipper
type Registrar struct {
	parser messages.MessageAddressesParser
}

// NewRegistrar allows to build a new Registrar instance
func NewRegistrar(parser messages.MessageAddressesParser) *Registrar {
	return &Registrar{
		parser: UniqueAddressesParser(parser),
	}
}

// BuildModules implements modules.Registrar
func (r *Registrar) BuildModules(ctx registrar.Context) jmodules.Modules {
	cdc := ctx.EncodingConfig.Codec
	db := database.Cast(ctx.Database)

	sources, err := types.BuildSources(ctx.JunoConfig.Node, ctx.EncodingConfig)
	if err != nil {
		panic(err)
	}

	actionsModule := actions.NewModule(ctx.JunoConfig, ctx.EncodingConfig)
	authModule := auth.NewModule(r.parser, cdc, db)
	bankModule := bank.NewModule(r.parser, sources.BankSource, cdc, db)
	consensusModule := consensus.NewModule(db)
	dailyRefetchModule := dailyrefetch.NewModule(ctx.Proxy, db)
	distrModule := distribution.NewModule(sources.DistrSource, cdc, db)
	feegrantModule := feegrant.NewModule(cdc, db)
	mintModule := mint.NewModule(sources.MintSource, cdc, db)
	slashingModule := slashing.NewModule(sources.SlashingSource, cdc, db)
	stakingModule := staking.NewModule(sources.StakingSource, cdc, db)
	govModule := gov.NewModule(sources.GovSource, authModule, distrModule, mintModule, slashingModule, stakingModule, cdc, db)
	upgradeModule := upgrade.NewModule(db, stakingModule)
	reputationModule := reputation.NewModule(cdc, db, sources.ReputationSource)
	identityModule := identity.NewModule(cdc, db, sources.IdentitySource, sources.RbacSource)
	marketplaceModule := marketplace.NewModule(cdc, db, sources.MarketplaceSource)
	productsModule := products.NewModule(cdc, db, sources.ProductsSource, sources.NftSource)
	rbacModule := rbac.NewModule(cdc, db)
	economicsModule := economics.NewModule(cdc, db, sources.EconomicsSource)
	tokenModule := token.NewModule(cdc, db, sources.TokenSource)

	return []jmodules.Module{
		messages.NewModule(r.parser, cdc, ctx.Database),
		telemetry.NewModule(ctx.JunoConfig),
		pruning.NewModule(ctx.JunoConfig, db, ctx.Logger),

		actionsModule,
		authModule,
		bankModule,
		consensusModule,
		dailyRefetchModule,
		distrModule,
		feegrantModule,
		govModule,
		mintModule,
		modules.NewModule(ctx.JunoConfig.Chain, db),
		pricefeed.NewModule(ctx.JunoConfig, cdc, db),
		slashingModule,
		stakingModule,
		upgradeModule,

		identityModule,
		rbacModule,
		reputationModule,
		tokenModule,
		productsModule,
		marketplaceModule,
		economicsModule,
	}
}
