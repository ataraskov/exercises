// Package prime provides functions to calculate prime numbers
package prime

import "errors"

var (
	ErrTooSmall = errors.New("too small number")
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, ErrTooSmall
	}
	count := 0
	for i := 0; ; i++ {
		if isPrime(i) {
			count++
			if count == n {
				return i, nil
			}
		}
	}

}

// isPrime checks if given number is prime
func isPrime(n int) bool {
	switch {
	case n < 2:
		return false
	case n == 2:
		return true
	default:
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
}
