// Pacakge luhn implements a simple Luhn algorithm
package luhn

import (
	"unicode"
)

// Valid checks if a given string is a valid Luhn number
func Valid(id string) bool {
	// convert runes to int slice
	nums := make([]int, 0)
	for _, r := range id {
		switch {
		case unicode.IsNumber(r):
			nums = append(nums, int(r-'0'))
		case unicode.IsSpace(r):
			continue
		default:
			return false
		}
	}

	// short ids are invalid by definition
	if len(nums) <= 1 {
		return false
	}

	// double each second digit (from the right)
	nums = Reverse(nums)
	for i, n := range nums {
		if (i+1)%2 == 0 {
			n += n
			if n > 9 {
				n -= 9
			}
			nums[i] = n
		}
	}

	// sum the digits
	sum := 0
	for _, n := range nums {
		sum += n
	}

	// if sum is not divisible by 10, it's not valid
	if sum%10 != 0 {
		return false
	}

	return true
}

// Reverse reverses an int slice
func Reverse(input []int) []int {
	var output []int
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output
}
