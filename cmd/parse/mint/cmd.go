package mint

import (
	"github.com/spf13/cobra"
	parsecmdtypes "github.com/villagelabsco/juno/v4/cmd/parse/types"
)

// NewMintCmd returns the Cobra command allowing to fix various things related to the x/mint module
func NewMintCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint",
		Short: "Fix things related to the x/mint module",
	}

	cmd.AddCommand(
		inflationCmd(parseConfig),
	)

	return cmd
}
