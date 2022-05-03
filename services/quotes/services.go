package quotes

import (
	"math/rand"
	"statham_quotes_server/models"
)

var quotes = []string{"I'm certainly not Tom Cruise or Brad Pitt."}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetQuote() models.Quote {
	return models.Quote{Text: quotes[rand.Intn(len(quotes))]}
}
