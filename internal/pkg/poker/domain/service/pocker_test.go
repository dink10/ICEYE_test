package service

import (
	"testing"
)

var testSuite = []struct {
	hand1  string
	hand2  string
	winner int
	title  string
}{
	{
		"AAAQQ",
		"QQAAA",
		0,
		"Tie",
	},
	{
		"53QQ2",
		"Q53Q2",
		0,
		"Tie",
	},
	{
		"53888",
		"88385",
		0,
		"Tie",
	},
	{
		"QQAAA",
		"AAAQQ",
		0,
		"Tie",
	},
	{
		"Q53Q2",
		"53QQ2",
		0,
		"Tie",
	},
	{
		"88385",
		"53888",
		0,
		"Tie",
	},
	{
		"AAAQQ",
		"QQQAA",
		1,
		"Hand 1",
	},
	{
		"Q53Q4",
		"53QQ2",
		1,
		"Hand 1",
	},
	{
		"53888",
		"88375",
		1,
		"Hand 1",
	},
	{
		"33337",
		"QQAAA",
		1,
		"Hand 1",
	},
	{
		"22333",
		"AAA58",
		1,
		"Hand 1",
	},
	{
		"33389",
		"AAKK4",
		1,
		"Hand 1",
	},
	{
		"44223",
		"AA892",
		1,
		"Hand 1",
	},
	{
		"22456",
		"AKQJT",
		1,
		"Hand 1",
	},
	{
		"99977",
		"77799",
		1,
		"Hand 1",
	},
	{
		"99922",
		"88866",
		1,
		"Hand 1",
	},
	{
		"9922A",
		"9922K",
		1,
		"Hand 1",
	},
	{
		"99975",
		"99965",
		1,
		"Hand 1",
	},
	{
		"99975",
		"99974",
		1,
		"Hand 1",
	},
	{
		"99752",
		"99652",
		1,
		"Hand 1",
	},

	{
		"99752",
		"99742",
		1,
		"Hand 1",
	},
	{
		"99753",
		"99752",
		1,
		"Hand 1",
	},
	{
		"QQQAA",
		"AAAQQ",
		2,
		"Hand 2",
	},
	{
		"53QQ2",
		"Q53Q4",
		2,
		"Hand 2",
	},
	{
		"88375",
		"53888",
		2,
		"Hand 2",
	},
	{
		"QQAAA",
		"33337",
		2,
		"Hand 2",
	},
	{
		"AAA58",
		"22333",
		2,
		"Hand 2",
	},
	{
		"AAKK4",
		"33389",
		2,
		"Hand 2",
	},
	{
		"AA892",
		"44223",
		2,
		"Hand 2",
	},
	{
		"AKQJT",
		"22456",
		2,
		"Hand 2",
	},
	{
		"77799",
		"99977",
		2,
		"Hand 2",
	},
	{
		"88866",
		"99922",
		2,
		"Hand 2",
	},
	{
		"9922K",
		"9922A",
		2,
		"Hand 2",
	},
	{
		"99965",
		"99975",
		2,
		"Hand 2",
	},
	{
		"99974",
		"99975",
		2,
		"Hand 2",
	},
	{
		"99652",
		"99752",
		2,
		"Hand 2",
	},
	{
		"99742",
		"99752",
		2,
		"Hand 2",
	},
	{
		"99752",
		"99753",
		2,
		"Hand 2",
	},
}

func TestWinner(t *testing.T) {
	srv := NewService()

	for _, ts := range testSuite {
		hand1 := make([]string, 0, 5)
		hand2 := make([]string, 0, 5)
		for i := range ts.hand1 {
			hand1 = append(hand1, string(ts.hand1[i]))
		}

		for i := range ts.hand1 {
			hand2 = append(hand2, string(ts.hand2[i]))
		}

		if w := srv.GetWinner(hand1, hand2); w != ts.winner {
			t.Fatalf(
				"[%s] [%s] Expected: %d, got: %d",
				ts.hand1, ts.hand2, ts.winner, w,
			)
		}
	}
}
