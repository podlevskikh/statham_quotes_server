package quotes

import (
	"github.com/podlevskikh/statham_quotes_server/models"
	"math/rand"
)

var quotes = []string{
	"I'm certainly not Tom Cruise or Brad Pitt.",
	"The best things in life sometimes happen spontaneously.",
	"I'm not as big a soccer fan as people might imagine, being British.",
	"I love to get behind the wheel and get competitive.",
	"I've come from nowhere, and I'm not shy to go back.",
	"Do 40 hard minutes, not an hour and a half of nonsense.",
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetQuote() models.Quote {
	return models.Quote{Text: quotes[rand.Intn(len(quotes))]}
}
