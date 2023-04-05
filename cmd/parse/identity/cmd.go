package bank

import (
	"github.com/spf13/cobra"
	parsecmdtypes "github.com/villagelabsco/juno/v4/cmd/parse/types"
)

// NewBankCmd returns the Cobra command allowing to fix various things related to the x/bank module
func NewIdentityCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identity",
		Short: "Fix things related to the idenitity module",
	}

	cmd.AddCommand(
		createNetworkCmd(parseConfig),
		createHumanCmd(parseConfig),
	)

	return cmd
}
