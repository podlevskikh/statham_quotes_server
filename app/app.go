package app

import (
	"context"
	"github.com/podlevskikh/statham_quotes_server/app/api"
	"github.com/podlevskikh/statham_quotes_server/services/pow"
	"github.com/podlevskikh/statham_quotes_server/services/quotes"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type App struct {
	logger *zerolog.Logger
}

func (a *App) Start(ctx context.Context, logger *zerolog.Logger) error {
	a.logger = logger

	quoteService := quotes.NewService()
	powService := pow.NewService()

	tcpAPI := api.NewTcpAPI(quoteService, powService, a.logger)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return tcpAPI.RunTCPServer(ctx)
	})
	return g.Wait()
}
