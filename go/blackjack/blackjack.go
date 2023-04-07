package blackjack

const (
	SPLIT   = "P"
	HIT     = "H"
	STAND   = "S"
	AUTOWIN = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
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
