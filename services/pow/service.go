package pow

import (
	"github.com/pkg/errors"
	"github.com/podlevskikh/statham_quotes_server/models/hashcash"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetChallenge(userInfo string) string {
	return hashcash.NewHashcash(userInfo).ToString()
}

func (s *Service) Validate(challenge, solution string) (bool, error) {
	hcChallenge, err := hashcash.Parse(challenge)
	if err != nil {
		return false, errors.Wrap(err, "challenge parse")
	}
	hcSolution, err := hashcash.Parse(solution)
	if err != nil {
		return false, errors.Wrap(err, "solution parse")
	}
	if hcChallenge.Ver != hcSolution.Ver ||
		hcChallenge.Bits != hcSolution.Bits ||
		!hcChallenge.Date.Equal(hcSolution.Date) ||
		hcChallenge.Resource != hcSolution.Resource ||
		hcChallenge.Rand != hcSolution.Rand {
		return false, errors.Wrap(err, "challenge solution match")
	}

	return hcSolution.IsValid(), nil
}
