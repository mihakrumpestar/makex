package configs

import (
	"context"
	"fmt"
	configs_deploy "makex/configs/deploy"
	configs_destroy "makex/configs/destroy"
	"makex/configs/flags"
	"makex/internal/helpers"
	"makex/pkg/orchestrators"
	"makex/pkg/secrets"

	"github.com/gookit/goutil/dump"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "makex",
	Short:   "An opinionated tool for streamlined deployment of containers with secrets",
	Version: "v0.0.1",
	Long: `makex is an opinionated tool for streamlined deployment of containers with secrets.
			Author: Miha Krumpestar
			License: GPLv3 License`,
}

func Execute(ctx context.Context, configFiles []string) error {
	rootCmd.PersistentFlags().StringVar(flags.Flags.Subfolder, "subfolder", ".", "subfolder where targets reside or the current directory")
	var orchestratorsLocal = new([]string)
	rootCmd.PersistentFlags().StringArrayVarP(orchestratorsLocal, "orchestrator", "o", []string{orchestrators.DockerCompose.String()}, "orchestrator plugins to use")
	var secretsLocal = new(string)
	rootCmd.PersistentFlags().StringVarP(secretsLocal, "secrets", "s", secrets.Sops.String(), "secrets plugin to use")
	rootCmd.PersistentFlags().StringVarP(flags.Flags.Target, "target", "t", "", "name of the project/subproject you want to do things to")
	rootCmd.PersistentFlags().BoolVarP(flags.Flags.MultipleTargets, "multipleTargets", "m", false, "if multiple targets in subdirectory (default is false)")
	rootCmd.PersistentFlags().StringVarP(flags.Flags.Environment, "environment", "e", "", "the environment to load (default is empty)")

	// Bind the current command's flags to Viper, this won't place Vipers values into Cobra flags automatically tho
	err := viper.BindPFlags(rootCmd.PersistentFlags())
	if err != nil {
		return fmt.Errorf("Error binding flags to Viper: %v\n", err)
	}

	err = initConfig(configFiles)
	if err != nil {
		return err
	}

	UpdateFlagsFromViper(orchestratorsLocal, secretsLocal)

	// Load subpath now that we have it in config
	subfolder := *flags.Flags.Subfolder
	multipleTargets := *flags.Flags.MultipleTargets
	if subfolder != "" && multipleTargets {
		if multipleTargets {
			subfolder += "/" + *flags.Flags.Target
		}
		configFiles, err = helpers.FindFilesFromBaseDir(subfolder, []string{""}, "makex", []string{"yaml", "yml"}, false)
		if err != nil {
			log.Fatal().Stack().Err(err).Msg("")
		}

		UpdateFlagsFromViper(orchestratorsLocal, secretsLocal)
	}

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		dump.P(flags.Flags)
	}

	rootCmd.AddCommand(configs_deploy.DeployCmd)
	rootCmd.AddCommand(configs_destroy.DestroyCmd)

	return rootCmd.Execute()
}

func initConfig(configFiles []string) error {
	if len(configFiles) != 0 {
		err := readAndMergeConfigFiles(configFiles)
		if err != nil {
			return err
		}

		err = viper.ReadInConfig()
		if err != nil {
			return fmt.Errorf("ReadInConfig: %s", err)
		}
	} else {
		log.Warn().Msg("no config files defined")
	}

	viper.SetEnvPrefix("MAKEX")
	viper.AutomaticEnv()

	return nil
}

func readAndMergeConfigFiles(configFiles []string) error {
	for _, configFile := range configFiles {
		viper.SetConfigFile(configFile)

		err := viper.MergeInConfig()
		if err != nil {
			return fmt.Errorf("Fatal error in reading config file %s: %s \n", configFile, err)
		}
	}

	return nil
}

func UpdateFlagsFromViper(orchestratorsLocal *[]string, secretsLocal *string) error {
	var err error

	rootCmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		// Only update the flags if they haven't been set by the user or stayed default
		if !f.Changed {
			val := viper.Get(f.Name)
			if val != nil {
				err = rootCmd.PersistentFlags().Set(f.Name, fmt.Sprintf("%v", val))
				if err != nil {
					return
				}
			}
		}
	})

	if err != nil {
		return err
	}

	// validate flags
	orchestratorsL, err := orchestrators.FromStringArray(*orchestratorsLocal)
	if err != nil {
		return err
	}
	flags.Flags.Orchestrator = &orchestratorsL

	secretsL, err := secrets.FromString(*secretsLocal)
	if err != nil {
		return err
	}
	flags.Flags.Secrets = &secretsL

	return nil
}
