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
	"People like to pigeonhole you.",
	"Looking good and feeling good go hand in hand. If you have a healthy lifestyle, your diet and nutrition are set, and you’re working out, you’re going to feel good.",
	"I’m enthusiastic and ambitious, and I work hard.",
	"Do 40 hard minutes, not an hour and a half of nonsense.",
	"There is something about yourself that you don’t know...",
	"Every sequel needs to be bigger and better.",
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetQuote() models.Quote {
	return models.Quote{Text: quotes[rand.Intn(len(quotes))]}
}
