package entity

import (
	"strings"
)

// Category - card value.
type Category byte

// Tiebreaker - for cards comparing.
type Tiebreaker []byte

// Counter - cards counter.
type Counter map[byte]int

// Rank - card rank.
type Rank struct {
	Category
	Tiebreaker []byte
}

const (
	// HighCard combination
	HighCard Category = iota
	// OnePair combination
	OnePair
	// TwoPairs combination
	TwoPairs
	// ThreeOfAKind combination
	ThreeOfAKind
	// FullHouse combination
	FullHouse
	// FourOfAKind combination
	FourOfAKind
)

// Order - for ranking by card value.
const Order = "23456789TJQKA"

// CardRank - card rank.
func CardRank(b byte) int {
	return strings.Index(Order, string(b))
}

// Has - Iterate through counts, and see if the number we want has a corresponding card value.
func (c Counter) Has(num int) bool {
	for _, count := range c {
		if count == num {
			return true
		}
	}

	return false
}

func (c Counter) getValue(num int) (values []byte) {
	for val, count := range c {
		if count == num {
			values = append(values, val)
		}
	}

	return
}
