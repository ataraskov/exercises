package main

// Problem: https://leetcode.com/problems/find-the-sum-of-encrypted-integers/

func encrypt(x int) int {
	max := 0
	digits := 0
	for x > 0 {
		digit := x % 10
		if digit > max {
			max = digit
		}
		x /= 10
		digits++
	}
	total := max
	for digits > 1 {
		total *= 10
		total += max
		digits--
	}
	return total
}

func sumOfEncryptedInt(nums []int) int {
	total := 0

	for _, num := range nums {
		total += encrypt(num)
	}
	return total
}
