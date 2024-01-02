package configs_deploy

import (
	"makex/configs/flags"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy containers according to the specified orchestrator and target",
	RunE: func(cmd *cobra.Command, args []string) error {

		// prefix := ".enc."

		log.Info().Str("Subfolder", *flags.Flags.Subfolder).Msg("")

		return nil
	},
}
