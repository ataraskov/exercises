package main

// Problem: https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-i/

func minOperations(nums []int, k int) int {
	n := 0
	for _, num := range nums {
		if num < k {
			n++
		}
	}
	return n
}
