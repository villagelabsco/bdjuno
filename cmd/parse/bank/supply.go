package bank

import (
	"fmt"

	modulestypes "github.com/villagelabsco/bdjuno/v3/modules/types"

	"github.com/spf13/cobra"
	parsecmdtypes "github.com/villagelabsco/juno/v4/cmd/parse/types"
	"github.com/villagelabsco/juno/v4/types/config"

	"github.com/villagelabsco/bdjuno/v3/database"
	"github.com/villagelabsco/bdjuno/v3/modules/bank"
)

// supplyCmd returns the Cobra command allowing to refresh x/bank total supply
func supplyCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
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
