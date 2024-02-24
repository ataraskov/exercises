package main

// Problem: https://leetcode.com/problems/single-number

// Memory complexity: O(1)

// test cases
// [2,2,1] => 1
// [1,2,2] => 1
// [2,1,2] => 1
func singleNumber(nums []int) int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	return xor
}
