package service

import (
	"fmt"
	"strings"

	"github.com/dink10/poker/internal/pkg/poker/domain/entity"
)

// Service implements parsing functionality.
type Service struct{}

// NewService provides parsing service.
func NewService() *Service {
	return &Service{}
}

// GetHandFromString builds hand from input.
func (s *Service) GetHandFromString(input string) (entity.Hand, error) {
	hand := strings.Split(input, "")

	if err := s.validateHand(hand); err != nil {
		return nil, err
	}

	return hand, nil
}

func (s *Service) validateHand(input []string) error {
	if len(input) != 5 {
		return fmt.Errorf("incorrect cards number in input %s, expected: 5, got %d", input, len(input))
	}

	for i := range input {
		// nolint
		if !strings.Contains(entity.Order, input[i]) {
			return fmt.Errorf("incorrect cards in input %s, allowed only %s", input, entity.Order)
		}
	}

	return nil
}

// GetWinner - return decision about winner.
func (s *Service) GetWinner(p1, p2 entity.Hand) int {
	p1rank, p2rank := p1.Rank(), p2.Rank()

	switch {
	case p1rank.Category > p2rank.Category:
		return 1
	case p1rank.Category < p2rank.Category:
		return 2
	case p1rank.Category == p2rank.Category:
		for i := len(p1rank.Tiebreaker) - 1; i >= 0; i-- {
			cr1 := entity.CardRank(p1rank.Tiebreaker[i])
			cr2 := entity.CardRank(p2rank.Tiebreaker[i])
			switch {
			case cr1 > cr2:
				return 1
			case cr1 < cr2:
				return 2
			}
		}

		for i := len(p1) - 1; i >= 0; i-- {
			if p1[i] > p2[i] {
				return 1
			}
			if p1[i] < p2[i] {
				return 2
			}
		}
	}

	return 0
}
