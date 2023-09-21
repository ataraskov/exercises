// Package yacht provides a solution for exercise Yacht on Exercism's Go Track.
package yacht

// Score returns a score for given dices within category.
func Score(dice []int, category string) int {
	count := make(map[int]int)
	for _, n := range dice {
		count[n]++
	}

	switch category {
	case "ones":
		return count[1] * 1
	case "twos":
		return count[2] * 2
	case "threes":
		return count[3] * 3
	case "fours":
		return count[4] * 4
	case "fives":
		return count[5] * 5
	case "sixes":
		return count[6] * 6
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
		return sum(dice)
	case "yacht":
		if len(count) == 1 {
			return 50
		}
	}
	return 0
}

// sum returns a sum of slice elements
func sum(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}
