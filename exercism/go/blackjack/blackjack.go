package blackjack

const (
	SPLIT   = "P"
	HIT     = "H"
	STAND   = "S"
	AUTOWIN = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cards := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}

	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1 := ParseCard(card1)
	c2 := ParseCard(card2)
	d := ParseCard(dealerCard)
	switch {
	case c1 == c2 && c1 == 11:
		return SPLIT
	case c1+c2 == 21 && d < 10:
		return AUTOWIN
	case c1+c2 == 21 && d > 9:
		return STAND
	case c1+c2 > 16 && c1+c2 < 21:
		return STAND
	case c1+c2 > 11 && c1+c2 < 17 && d < 7:
		return STAND
	case c1+c2 > 11 && c1+c2 < 17 && d > 6:
		return HIT
	case c1+c2 < 12:
		return HIT
	}

	// should not reach this one, something wrong
	return "?"
}
