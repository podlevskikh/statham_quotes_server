package main

import (
	"context"
	"github.com/rs/zerolog"
	"os"
	"statham_quotes_server/app"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().With().Str("source", "statham_quotes_server").Logger()
	ctx := context.Background()
	a := app.App{}
	if err := a.Start(ctx, &logger); err != nil {
		logger.Fatal().Err(err).Msg("server start")
	}
}
