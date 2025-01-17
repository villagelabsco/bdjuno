package distribution

import (
	"fmt"

	"github.com/spf13/cobra"
	parsecmdtypes "github.com/villagelabsco/juno/v4/cmd/parse/types"
	"github.com/villagelabsco/juno/v4/types/config"

	"github.com/villagelabsco/bdjuno/v3/database"
	"github.com/villagelabsco/bdjuno/v3/modules/distribution"
	modulestypes "github.com/villagelabsco/bdjuno/v3/modules/types"
)

// communityPoolCmd returns the Cobra command allowing to refresh community pool
func communityPoolCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "community-pool",
		Short: "Refresh community pool",
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

			// Build distribution module
			distrModule := distribution.NewModule(sources.DistrSource, parseCtx.EncodingConfig.Codec, db)

			err = distrModule.GetLatestCommunityPool()
			if err != nil {
				return fmt.Errorf("error while updating community pool: %s", err)
			}

			return nil
		},
	}
}
