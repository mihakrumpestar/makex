package configs_destroy

import "github.com/spf13/cobra"

var DestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroys containers according to the specified orchestrator and target",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}
