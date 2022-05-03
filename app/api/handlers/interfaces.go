package handlers

import "github.com/podlevskikh/statham_quotes_server/models"

type QuotesService interface {
	GetQuote() models.Quote
}

type POWService interface {
	GetChallenge(userInfo string) string
	Validate(challenge, solution string) (bool, error)
}
