package main

import (
	"context"
	"github.com/podlevskikh/statham_quotes_server/app"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().With().Str("source", "statham_quotes_server").Logger()
	ctx := context.Background()
	a := app.App{}
	if err := a.Start(ctx, &logger); err != nil {
		logger.Fatal().Err(err).Msg("server start")
	}
}
