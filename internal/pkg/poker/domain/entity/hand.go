package entity

import (
	"sort"
)

const (
	// HandsCount - count of players
	HandsCount = 2
)

// Hand - cards of player
type Hand []string

func (h Hand) Len() int {
	return 5
}

func (h Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hand) Less(i, j int) bool {
	return CardRank(h[i][0]) > CardRank(h[j][0])
}

// Tiebreaker - To use the whole hand's values into a tiebreaker list
func (h Hand) Tiebreaker() []byte {
	return []byte{h[0][0], h[1][0], h[2][0], h[3][0], h[4][0]}
}

// Rank - had's rank
func (h Hand) Rank() Rank {
	sort.Sort(h)

	// Look for fours, threes, and  pairs, in that order, and create tiebreaker
	CardCounter := make(Counter)
	for _, card := range h {
		CardCounter[card[0]]++
	}
	var tie Tiebreaker
	for _, n := range []int{4, 3, 2} {
		if CardCounter.Has(n) {
			switch n {
			case 4:
				// Both players can't have the same 4, so it's the only tiebreaker.
				return Rank{FourOfAKind, CardCounter.getValue(4)}
			case 3:
				// Both players can't have the same 3, so it's the only tiebreaker.
				tie = CardCounter.getValue(3)
				if CardCounter.Has(2) {
					return Rank{FullHouse, tie}
				}
				return Rank{ThreeOfAKind, tie}
			case 2:
				if len(CardCounter) == 3 {
					// 2 pairs, each of which can break tie, but first put the kicker in the lowest position.
					tie = CardCounter.getValue(1)
					pairs := CardCounter.getValue(2)
					// Sort pairs
					if pairs[0] > pairs[1] {
						pairs[0], pairs[1] = pairs[1], pairs[0]
					}
					for _, v := range pairs {
						tie = append(tie, v)
					}
					return Rank{TwoPairs, tie}
				}
				// len == 4, or 1 pair
				tie = CardCounter.getValue(1)
				sort.Slice(tie, func(i, j int) bool {
					return tie[i] < tie[j]
				})
				// One pair
				tie = append(tie, CardCounter.getValue(2)[0])
				return Rank{OnePair, tie} // Only one card left.
			}
		}
	}

	return Rank{HighCard, h.Tiebreaker()}
}
