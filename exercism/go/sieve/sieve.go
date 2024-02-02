// Package sieve implements a simple sieve of Eratosthenes
package sieve

// Sieve returns a slice of prime numbers up to limit
func Sieve(limit int) []int {
	composites := make([]bool, limit+1)
	primes := []int{}

	for num := 2; num <= limit; num++ {
		if !composites[num] {
			primes = append(primes, num)
		}
		for i := num + num; i <= limit; i += num {
			composites[i] = true
		}
	}

	return primes
}
