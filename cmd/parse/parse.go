package parse

import (
	"github.com/spf13/cobra"
	parse "github.com/villagelabsco/juno/v4/cmd/parse/types"

	parseblocks "github.com/villagelabsco/juno/v4/cmd/parse/blocks"

	parsegenesis "github.com/villagelabsco/juno/v4/cmd/parse/genesis"

	parseauth "github.com/villagelabsco/bdjuno/v3/cmd/parse/auth"
	parsebank "github.com/villagelabsco/bdjuno/v3/cmd/parse/bank"
	parsedistribution "github.com/villagelabsco/bdjuno/v3/cmd/parse/distribution"
	parsefeegrant "github.com/villagelabsco/bdjuno/v3/cmd/parse/feegrant"
	parsegov "github.com/villagelabsco/bdjuno/v3/cmd/parse/gov"
	parseidentity "github.com/villagelabsco/bdjuno/v3/cmd/parse/identity"
	parsemint "github.com/villagelabsco/bdjuno/v3/cmd/parse/mint"
	parsepricefeed "github.com/villagelabsco/bdjuno/v3/cmd/parse/pricefeed"
	parsestaking "github.com/villagelabsco/bdjuno/v3/cmd/parse/staking"
	parsetransaction "github.com/villagelabsco/juno/v4/cmd/parse/transactions"
)

// NewParseCmd returns the Cobra command allowing to parse some chain data without having to re-sync the whole database
func NewParseCmd(parseCfg *parse.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "parse",
		Short:             "Parse some data without the need to re-syncing the whole database from scratch",
		PersistentPreRunE: runPersistentPreRuns(parse.ReadConfigPreRunE(parseCfg)),
	}

	cmd.AddCommand(
		parseauth.NewAuthCmd(parseCfg),
		parsebank.NewBankCmd(parseCfg),
		parseblocks.NewBlocksCmd(parseCfg),
		parsedistribution.NewDistributionCmd(parseCfg),
		parsefeegrant.NewFeegrantCmd(parseCfg),
		parsegenesis.NewGenesisCmd(parseCfg),
		parsegov.NewGovCmd(parseCfg),
		parsemint.NewMintCmd(parseCfg),
		parsepricefeed.NewPricefeedCmd(parseCfg),
		parsestaking.NewStakingCmd(parseCfg),
		parsetransaction.NewTransactionsCmd(parseCfg),
		parseidentity.NewIdentityCmd(parseCfg),
	)

	return cmd
}

func runPersistentPreRuns(preRun func(_ *cobra.Command, _ []string) error) func(_ *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if root := cmd.Root(); root != nil {
			if root.PersistentPreRunE != nil {
				err := root.PersistentPreRunE(root, args)
				if err != nil {
					return err
				}
			}
		}

		return preRun(cmd, args)
	}
}
