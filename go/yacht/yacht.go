// Package yacht provides a solution for exercise Yacht on Exercism's Go Track.
package yacht

// Score returns a score for given dices within category.
func Score(dice []int, category string) int {
	count := make(map[int]int)
	sum := 0
	for _, n := range dice {
		count[n]++
		sum += n
	}

	cat2num := map[string]int{
		"ones":   1,
		"twos":   2,
		"threes": 3,
		"fours":  4,
		"fives":  5,
		"sixes":  6,
	}

	switch category {
	case "ones", "twos", "threes", "fours", "fives", "sixes":
		n := cat2num[category]
		return count[n] * n
	case "full house":
		if len(count) == 2 {
			sum := 0
			for k, v := range count {
				if v == 3 || v == 2 {
					sum += k * v
				}
			}
			return sum
		}
	case "four of a kind":
		for k, v := range count {
			if v >= 4 {
				return k * 4
			}
		}
	case "little straight":
		if len(count) == 5 && count[6] == 0 {
			return 30
		}
	case "big straight":
		if len(count) == 5 && count[1] == 0 {
			return 30
		}
	case "choice":
		return sum
	case "yacht":
		if len(count) == 1 {
			return 50
		}
	}
	return 0
}
