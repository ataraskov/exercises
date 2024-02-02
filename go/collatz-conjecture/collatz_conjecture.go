package collatzconjecture

import "errors"

// CollatzConjecture calculates the number of steps required to reach 1 for 3x+1 problem
// https://en.wikipedia.org/wiki/Collatz_conjecture
func CollatzConjecture(n int) (int, error) {
	steps := 0
	if n < 1 {
		return 0, errors.New("invalid number")
	}

	for ; n > 1; steps++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}
	return steps, nil
}
