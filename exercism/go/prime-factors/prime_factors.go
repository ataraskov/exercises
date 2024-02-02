// Package prime provides functions to calculate prime factors
package prime

// Factors calculates prime factors of a number
func Factors(n int64) []int64 {
	res := []int64{}
	var factor int64
	for factor = 2; n > 1; {
		if n%factor == 0 {
			n /= factor
			res = append(res, factor)
		} else {
			factor++
		}
	}
	return res
}
