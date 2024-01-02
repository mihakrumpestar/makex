package cmd

import (
	"context"
	"makex/configs"
	"makex/internal/helpers"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func PrepareAndExecute(ctx context.Context) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	configFiles, err := helpers.FindFilesFromBaseDir("", []string{""}, "makex", []string{"yaml", "yml"}, false)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("")
	}

	if err := configs.Execute(ctx, configFiles); err != nil {
		log.Fatal().Stack().Err(err).Msg("")
	}
}
