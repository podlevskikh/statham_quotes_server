package handlers

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog"
	"net"
	"strings"
)

type Quote struct {
	quotesService QuotesService
	powService    POWService
	logger        *zerolog.Logger
}

func NewQuote(quotesService QuotesService, powService POWService, logger *zerolog.Logger) *Quote {
	return &Quote{
		quotesService: quotesService,
		powService:    powService,
		logger:        logger,
	}
}

func (h *Quote) HandleRequest(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error close connection:", err.Error())
		}
	}()

	challenge := h.powService.GetChallenge(strings.Split(conn.RemoteAddr().String(),":")[0])

	_, err := conn.Write([]byte(challenge))
	if err != nil {
		h.logger.Err(err).Msg("send challenge")
		return
	}

	reader := bufio.NewReader(conn)
	solution, err := reader.ReadString('\n')
	if err != nil {
		h.logger.Err(err).Msg("read solution")
		return
	}

	valid, err := h.powService.Validate(challenge, strings.Trim(solution, "\n"))
	if err != nil {
		h.logger.Err(err).Msg("validate error")
		return
	}

	if !valid {
		h.logger.Err(err).Msg("solution incorrect")
		return
	}

	_, err = conn.Write([]byte(h.quotesService.GetQuote().Text))
	if err != nil {
		h.logger.Err(err).Msg("write quote")
		return
	}
}
