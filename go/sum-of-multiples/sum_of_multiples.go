// Package summultiples implements a solution to the Sum of Multiples exercise at exercism.io
package summultiples

// SumMultiples returns the sum of all multiples of divisors up to but not including limit
func SumMultiples(limit int, divisors ...int) int {
	if limit < 1 {
		return 0
	}

	multiplies := make([]bool, limit)
	for _, divisor := range divisors {
		if divisor < 1 {
			continue
		}
		for n := divisor; n < limit; n += divisor {
			multiplies[n] = true
		}
	}

	res := 0
	for i := 1; i < limit; i++ {
		if multiplies[i] {
			res += i
		}
	}
	return res
}
