package bank

import (
	"encoding/json"
	"fmt"

	"github.com/villagelabsco/bdjuno/v3/modules/identity"
	modulestypes "github.com/villagelabsco/bdjuno/v3/modules/types"

	"github.com/spf13/cobra"
	parsecmdtypes "github.com/villagelabsco/juno/v4/cmd/parse/types"
	"github.com/villagelabsco/juno/v4/types/config"

	"github.com/villagelabsco/bdjuno/v3/database"
	"github.com/villagelabsco/bdjuno/v3/modules/bank"
	"github.com/villagelabsco/bdjuno/v3/utils"
)

// supplyCmd returns the Cobra command allowing to refresh x/bank total supply
func createNetworkCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "network",
		Short: "Fill in missing networks",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			sources, err := modulestypes.BuildSources(config.Cfg.Node, parseCtx.EncodingConfig)
			if err != nil {
				return err
			}

			genesis, err := utils.ReadGenesis(config.Cfg, parseCtx.Node)

			var appState map[string]json.RawMessage
			if err := json.Unmarshal(genesis.AppState, &appState); err != nil {
				return fmt.Errorf("error unmarshalling genesis doc: %s", err)
			}

			networks, err := identity.GetNetworksFromGenesis(appState, parseCtx.EncodingConfig.Codec)

			providers, err := identity.GetProvidersFromGenesis(appState, parseCtx.EncodingConfig.Codec)
			// Get the database
			db := database.Cast(parseCtx.Database)

			// Build bank module
			identityModule := identity.NewModule(parseCtx.EncodingConfig.Codec, db, sources.IdentitySource, sources.RbacSource, sources.FeeGrantSource)

			identityModule.UpdateProviders(providers)
			identityModule.UpdateNetworks(networks)

			if err != nil {
				return fmt.Errorf("error while getting latest bank supply: %s", err)
			}

			return nil
		},
	}
}

func createHumanCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "supply",
		Short: "Refresh total supply",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			sources, err := modulestypes.BuildSources(config.Cfg.Node, parseCtx.EncodingConfig)
			if err != nil {
				return err
			}

			// Get the database
			db := database.Cast(parseCtx.Database)

			// Build bank module
			bankModule := bank.NewModule(nil, sources.BankSource, parseCtx.EncodingConfig.Codec, db)

			err = bankModule.UpdateSupply()
			if err != nil {
				return fmt.Errorf("error while getting latest bank supply: %s", err)
			}

			return nil
		},
	}
}
