package api

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"net"
	"statham_quotes_server/app/api/handlers"
	powService "statham_quotes_server/services/pow"
	quotesService "statham_quotes_server/services/quotes"
)

type TcpAPI struct {
	powService    *powService.Service
	quotesService *quotesService.Service
	logger        *zerolog.Logger
}

func NewTcpAPI(quotesService *quotesService.Service, powService *powService.Service, logger *zerolog.Logger) *TcpAPI {
	return &TcpAPI{quotesService: quotesService, powService: powService, logger: logger}
}

func (a *TcpAPI) RunTCPServer(ctx context.Context) error {
	l, err := net.Listen("tcp", "localhost:3333")
	if err != nil {
		return errors.Wrap(err, "listen")
	}
	defer func() {
		err = l.Close()
		if err != nil {
			a.logger.Err(err).Msg("close listener")
		}
	}()
	a.logger.Info().Msg("listen localhost:3333")
	h := handlers.NewQuote(a.quotesService, a.powService, a.logger)
	for {
		conn, err := l.Accept()
		if err != nil {
			return errors.Wrap(err, "accepting")
		}

		go h.HandleRequest(conn)
	}
}
