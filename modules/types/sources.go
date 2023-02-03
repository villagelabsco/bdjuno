package types

import (
	"fmt"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	econsource "github.com/villagelabsco/bdjuno/v3/modules/economics/source"
	identitysource "github.com/villagelabsco/bdjuno/v3/modules/identity/source"
	remotenftsource "github.com/villagelabsco/bdjuno/v3/modules/nft/source/remote"
	remotetokensource "github.com/villagelabsco/bdjuno/v3/modules/token/source/remote"
	econtypes "github.com/villagelabsco/villaged/x/economics/types"
	identitytypes "github.com/villagelabsco/villaged/x/identity/types"
	marketplacetypes "github.com/villagelabsco/villaged/x/marketplace/types"
	productstypes "github.com/villagelabsco/villaged/x/products/types"
	rbactypes "github.com/villagelabsco/villaged/x/rbac/types"
	reputationtypes "github.com/villagelabsco/villaged/x/reputation/types"
	tokentypes "github.com/villagelabsco/villaged/x/token/types"
	"os"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	v1beta1govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/villagelabsco/juno/v4/node/local"
	"github.com/villagelabsco/juno/v4/node/remote"

	nodeconfig "github.com/villagelabsco/juno/v4/node/config"

	banksource "github.com/villagelabsco/bdjuno/v3/modules/bank/source"
	localbanksource "github.com/villagelabsco/bdjuno/v3/modules/bank/source/local"
	remotebanksource "github.com/villagelabsco/bdjuno/v3/modules/bank/source/remote"
	distrsource "github.com/villagelabsco/bdjuno/v3/modules/distribution/source"
	localdistrsource "github.com/villagelabsco/bdjuno/v3/modules/distribution/source/local"
	remotedistrsource "github.com/villagelabsco/bdjuno/v3/modules/distribution/source/remote"
	remoteeconsource "github.com/villagelabsco/bdjuno/v3/modules/economics/source/remote"
	govsource "github.com/villagelabsco/bdjuno/v3/modules/gov/source"
	remotegovsource "github.com/villagelabsco/bdjuno/v3/modules/gov/source/remote"
	remoteidentitysource "github.com/villagelabsco/bdjuno/v3/modules/identity/source/remote"
	marketplacesource "github.com/villagelabsco/bdjuno/v3/modules/marketplace/source"
	remotemarketplacesource "github.com/villagelabsco/bdjuno/v3/modules/marketplace/source/remote"
	mintsource "github.com/villagelabsco/bdjuno/v3/modules/mint/source"
	localmintsource "github.com/villagelabsco/bdjuno/v3/modules/mint/source/local"
	remotemintsource "github.com/villagelabsco/bdjuno/v3/modules/mint/source/remote"
	nftsource "github.com/villagelabsco/bdjuno/v3/modules/nft/source"
	productssource "github.com/villagelabsco/bdjuno/v3/modules/products/source"
	remoteproductssource "github.com/villagelabsco/bdjuno/v3/modules/products/source/remote"
	rbacsource "github.com/villagelabsco/bdjuno/v3/modules/rbac/source"
	remoterbacsource "github.com/villagelabsco/bdjuno/v3/modules/rbac/source/remote"
	reputationsource "github.com/villagelabsco/bdjuno/v3/modules/reputation/source"
	remotereputationsource "github.com/villagelabsco/bdjuno/v3/modules/reputation/source/remote"
	slashingsource "github.com/villagelabsco/bdjuno/v3/modules/slashing/source"
	localslashingsource "github.com/villagelabsco/bdjuno/v3/modules/slashing/source/local"
	remoteslashingsource "github.com/villagelabsco/bdjuno/v3/modules/slashing/source/remote"
	stakingsource "github.com/villagelabsco/bdjuno/v3/modules/staking/source"
	localstakingsource "github.com/villagelabsco/bdjuno/v3/modules/staking/source/local"
	remotestakingsource "github.com/villagelabsco/bdjuno/v3/modules/staking/source/remote"
	tokensource "github.com/villagelabsco/bdjuno/v3/modules/token/source"
)

type Sources struct {
	BankSource     banksource.Source
	DistrSource    distrsource.Source
	GovSource      govsource.Source
	MintSource     mintsource.Source
	SlashingSource slashingsource.Source
	StakingSource  stakingsource.Source
	NftSource      nftsource.Source

	IdentitySource    identitysource.Source
	RbacSource        rbacsource.Source
	ReputationSource  reputationsource.Source
	ProductsSource    productssource.Source
	MarketplaceSource marketplacesource.Source
	EconomicsSource   econsource.Source
	TokenSource       tokensource.Source
}

func BuildSources(nodeCfg nodeconfig.Config, encodingConfig *params.EncodingConfig) (*Sources, error) {
	switch cfg := nodeCfg.Details.(type) {
	case *remote.Details:
		return buildRemoteSources(cfg)
	case *local.Details:
		return buildLocalSources(cfg, encodingConfig)

	default:
		return nil, fmt.Errorf("invalid configuration type: %T", cfg)
	}
}

func buildLocalSources(cfg *local.Details, encodingConfig *params.EncodingConfig) (*Sources, error) {
	source, err := local.NewSource(cfg.Home, encodingConfig)
	if err != nil {
		return nil, err
	}

	sapp := simapp.NewSimApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), source.StoreDB, nil, true, map[int64]bool{},
		cfg.Home, 0, simapp.MakeTestEncodingConfig(), simapp.EmptyAppOptions{},
	)

	sources := &Sources{
		BankSource:  localbanksource.NewSource(source, banktypes.QueryServer(sapp.BankKeeper)),
		DistrSource: localdistrsource.NewSource(source, distrtypes.QueryServer(sapp.DistrKeeper)),
		//GovSource:      localgovsource.NewSource(source, v1beta1govtypes.QueryServer(sapp.GovKeeper)),
		MintSource:     localmintsource.NewSource(source, minttypes.QueryServer(sapp.MintKeeper)),
		SlashingSource: localslashingsource.NewSource(source, slashingtypes.QueryServer(sapp.SlashingKeeper)),
		StakingSource:  localstakingsource.NewSource(source, stakingkeeper.Querier{Keeper: sapp.StakingKeeper}),
		//NftSource:         localnftsource.NewSource(source, nfttypes.QueryServer(sapp.NFTKeeper)),
		//KycSource:         localkycsource.NewSource(source, kyctypes.QueryServer(sapp.KycKeeper)),
		//MarketplaceSource: localmarketplacesource.NewSource(source, marketplacetypes.QueryServer(sapp.MarketplaceKeeper)),
		//ProductsSource:    localproductssource.NewSource(source, productstypes.QueryServer(sapp.ProductsKeeper)),
		//RbacSource:        localrbacsource.NewSource(source, rbactypes.QueryServer(sapp.RbacKeeper)),
		//ReputationSource:  localreputationsource.NewSource(source, reputationtypes.QueryServer(sapp.ReputationKeeper)),
		//VillageSource:     localvillagesource.NewSource(source, villagetypes.QueryServer(sapp.VillageKeeper)),
	}

	// Mount and initialize the stores
	err = source.MountKVStores(sapp, "keys")
	if err != nil {
		return nil, err
	}

	err = source.MountTransientStores(sapp, "tkeys")
	if err != nil {
		return nil, err
	}

	err = source.MountMemoryStores(sapp, "memKeys")
	if err != nil {
		return nil, err
	}

	err = source.InitStores()
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func buildRemoteSources(cfg *remote.Details) (*Sources, error) {
	source, err := remote.NewSource(cfg.GRPC)
	if err != nil {
		return nil, fmt.Errorf("error while creating remote source: %s", err)
	}

	return &Sources{
		BankSource:        remotebanksource.NewSource(source, banktypes.NewQueryClient(source.GrpcConn)),
		DistrSource:       remotedistrsource.NewSource(source, distrtypes.NewQueryClient(source.GrpcConn)),
		GovSource:         remotegovsource.NewSource(source, v1beta1govtypes.NewQueryClient(source.GrpcConn)),
		MintSource:        remotemintsource.NewSource(source, minttypes.NewQueryClient(source.GrpcConn)),
		SlashingSource:    remoteslashingsource.NewSource(source, slashingtypes.NewQueryClient(source.GrpcConn)),
		StakingSource:     remotestakingsource.NewSource(source, stakingtypes.NewQueryClient(source.GrpcConn)),
		NftSource:         remotenftsource.NewSource(source, nfttypes.NewQueryClient(source.GrpcConn)),
		IdentitySource:    remoteidentitysource.NewSource(source, identitytypes.NewQueryClient(source.GrpcConn)),
		RbacSource:        remoterbacsource.NewSource(source, rbactypes.NewQueryClient(source.GrpcConn)),
		ReputationSource:  remotereputationsource.NewSource(source, reputationtypes.NewQueryClient(source.GrpcConn)),
		ProductsSource:    remoteproductssource.NewSource(source, productstypes.NewQueryClient(source.GrpcConn)),
		MarketplaceSource: remotemarketplacesource.NewSource(source, marketplacetypes.NewQueryClient(source.GrpcConn)),
		EconomicsSource:   remoteeconsource.NewSource(source, econtypes.NewQueryClient(source.GrpcConn)),
		TokenSource:       remotetokensource.NewSource(source, tokentypes.NewQueryClient(source.GrpcConn)),
	}, nil
}
