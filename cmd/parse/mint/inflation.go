package mint

import (
	"fmt"

	"github.com/spf13/cobra"
	parsecmdtypes "github.com/villagelabsco/juno/v4/cmd/parse/types"
	"github.com/villagelabsco/juno/v4/types/config"

	"github.com/villagelabsco/bdjuno/v3/database"
	"github.com/villagelabsco/bdjuno/v3/modules/mint"
	modulestypes "github.com/villagelabsco/bdjuno/v3/modules/types"
)

// inflationCmd returns the Cobra command allowing to refresh x/mint inflation
func inflationCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "inflation",
		Short: "Refresh inflation",
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

			// Build mint module
			mintModule := mint.NewModule(sources.MintSource, parseCtx.EncodingConfig.Codec, db)

			err = mintModule.UpdateInflation()
			if err != nil {
				return fmt.Errorf("error while updating inflation: %s", err)
			}

			return nil
		},
	}
}
