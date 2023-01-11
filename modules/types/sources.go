package types

import (
	"fmt"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	remotenftsource "github.com/forbole/bdjuno/v3/modules/nft/source/remote"
	kyctypes "github.com/villagelabs/villaged/x/kyc/types"
	marketplacetypes "github.com/villagelabs/villaged/x/marketplace/types"
	productstypes "github.com/villagelabs/villaged/x/products/types"
	rbactypes "github.com/villagelabs/villaged/x/rbac/types"
	reputationtypes "github.com/villagelabs/villaged/x/reputation/types"
	villagetypes "github.com/villagelabs/villaged/x/village/types"
	"os"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/forbole/juno/v3/node/remote"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/forbole/juno/v3/node/local"

	nodeconfig "github.com/forbole/juno/v3/node/config"

	banksource "github.com/forbole/bdjuno/v3/modules/bank/source"
	localbanksource "github.com/forbole/bdjuno/v3/modules/bank/source/local"
	remotebanksource "github.com/forbole/bdjuno/v3/modules/bank/source/remote"
	distrsource "github.com/forbole/bdjuno/v3/modules/distribution/source"
	localdistrsource "github.com/forbole/bdjuno/v3/modules/distribution/source/local"
	remotedistrsource "github.com/forbole/bdjuno/v3/modules/distribution/source/remote"
	govsource "github.com/forbole/bdjuno/v3/modules/gov/source"
	localgovsource "github.com/forbole/bdjuno/v3/modules/gov/source/local"
	remotegovsource "github.com/forbole/bdjuno/v3/modules/gov/source/remote"
	kycsource "github.com/forbole/bdjuno/v3/modules/kyc/source"
	marketplacesource "github.com/forbole/bdjuno/v3/modules/marketplace/source"
	mintsource "github.com/forbole/bdjuno/v3/modules/mint/source"
	localmintsource "github.com/forbole/bdjuno/v3/modules/mint/source/local"
	remotemintsource "github.com/forbole/bdjuno/v3/modules/mint/source/remote"
	nftsource "github.com/forbole/bdjuno/v3/modules/nft/source"
	localnftsource "github.com/forbole/bdjuno/v3/modules/nft/source/local"
	productssource "github.com/forbole/bdjuno/v3/modules/products/source"
	rbacsource "github.com/forbole/bdjuno/v3/modules/rbac/source"
	reputationsource "github.com/forbole/bdjuno/v3/modules/reputation/source"
	slashingsource "github.com/forbole/bdjuno/v3/modules/slashing/source"
	localslashingsource "github.com/forbole/bdjuno/v3/modules/slashing/source/local"
	remoteslashingsource "github.com/forbole/bdjuno/v3/modules/slashing/source/remote"
	stakingsource "github.com/forbole/bdjuno/v3/modules/staking/source"
	localstakingsource "github.com/forbole/bdjuno/v3/modules/staking/source/local"
	remotestakingsource "github.com/forbole/bdjuno/v3/modules/staking/source/remote"
	villagesource "github.com/forbole/bdjuno/v3/modules/village/source"

	remotekycsource "github.com/forbole/bdjuno/v3/modules/kyc/source/remote"
	remotemarketplacesource "github.com/forbole/bdjuno/v3/modules/marketplace/source/remote"
	remoteproductssource "github.com/forbole/bdjuno/v3/modules/products/source/remote"
	remoterbacsource "github.com/forbole/bdjuno/v3/modules/rbac/source/remote"
	remotereputationsource "github.com/forbole/bdjuno/v3/modules/reputation/source/remote"
	remotevillagesource "github.com/forbole/bdjuno/v3/modules/village/source/remote"

	localkycsource "github.com/forbole/bdjuno/v3/modules/kyc/source/local"
	localmarketplacesource "github.com/forbole/bdjuno/v3/modules/marketplace/source/local"
	localproductssource "github.com/forbole/bdjuno/v3/modules/products/source/local"
	localrbacsource "github.com/forbole/bdjuno/v3/modules/rbac/source/local"
	localreputationsource "github.com/forbole/bdjuno/v3/modules/reputation/source/local"
	localvillagesource "github.com/forbole/bdjuno/v3/modules/village/source/local"
)

type Sources struct {
	BankSource     banksource.Source
	DistrSource    distrsource.Source
	GovSource      govsource.Source
	MintSource     mintsource.Source
	SlashingSource slashingsource.Source
	StakingSource  stakingsource.Source

	KycSource         kycsource.Source
	MarketplaceSource marketplacesource.Source
	ProductsSource    productssource.Source
	RbacSource        rbacsource.Source
	ReputationSource  reputationsource.Source
	VillageSource     villagesource.Source
	NftSource         nftsource.Source
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

	app := simapp.NewSimApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), source.StoreDB, nil, true, map[int64]bool{},
		cfg.Home, 0, simapp.MakeTestEncodingConfig(), simapp.EmptyAppOptions{},
	)

	sources := &Sources{
		BankSource:        localbanksource.NewSource(source, banktypes.QueryServer(app.BankKeeper)),
		DistrSource:       localdistrsource.NewSource(source, distrtypes.QueryServer(app.DistrKeeper)),
		GovSource:         localgovsource.NewSource(source, govtypes.QueryServer(app.GovKeeper)),
		MintSource:        localmintsource.NewSource(source, minttypes.QueryServer(app.MintKeeper)),
		SlashingSource:    localslashingsource.NewSource(source, slashingtypes.QueryServer(app.SlashingKeeper)),
		StakingSource:     localstakingsource.NewSource(source, stakingkeeper.Querier{Keeper: app.StakingKeeper}),
		NftSource:         localnftsource.NewSource(source, nfttypes.QueryServer(app.NftKeeper)),
		KycSource:         localkycsource.NewSource(source, kyctypes.QueryServer(app.KycKeeper)),
		MarketplaceSource: localmarketplacesource.NewSource(source, marketplacetypes.QueryServer(app.MarketplaceKeeper)),
		ProductsSource:    localproductssource.NewSource(source, productstypes.QueryServer(app.ProductsKeeper)),
		RbacSource:        localrbacsource.NewSource(source, rbactypes.QueryServer(app.RbacKeeper)),
		ReputationSource:  localreputationsource.NewSource(source, reputationtypes.QueryServer(app.ReputationKeeper)),
		VillageSource:     localvillagesource.NewSource(source, villagetypes.QueryServer(app.VillageKeeper)),
	}

	// Mount and initialize the stores
	err = source.MountKVStores(app, "keys")
	if err != nil {
		return nil, err
	}

	err = source.MountTransientStores(app, "tkeys")
	if err != nil {
		return nil, err
	}

	err = source.MountMemoryStores(app, "memKeys")
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
		GovSource:         remotegovsource.NewSource(source, govtypes.NewQueryClient(source.GrpcConn)),
		MintSource:        remotemintsource.NewSource(source, minttypes.NewQueryClient(source.GrpcConn)),
		SlashingSource:    remoteslashingsource.NewSource(source, slashingtypes.NewQueryClient(source.GrpcConn)),
		StakingSource:     remotestakingsource.NewSource(source, stakingtypes.NewQueryClient(source.GrpcConn)),
		NftSource:         remotenftsource.NewSource(source, nfttypes.NewQueryClient(source.GrpcConn)),
		KycSource:         remotekycsource.NewSource(source, kyctypes.NewQueryClient(source.GrpcConn)),
		MarketplaceSource: remotemarketplacesource.NewSource(source, marketplacetypes.NewQueryClient(source.GrpcConn)),
		ProductsSource:    remoteproductssource.NewSource(source, productstypes.NewQueryClient(source.GrpcConn)),
		RbacSource:        remoterbacsource.NewSource(source, rbactypes.NewQueryClient(source.GrpcConn)),
		ReputationSource:  remotereputationsource.NewSource(source, reputationtypes.NewQueryClient(source.GrpcConn)),
		VillageSource:     remotevillagesource.NewSource(source, villagetypes.NewQueryClient(source.GrpcConn)),
	}, nil
}
