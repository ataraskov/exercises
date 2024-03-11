package main

// Problem: https://leetcode.com/problems/trapping-rain-water/

func trap(height []int) int {
	// brute force
	// -------------------------------------------------------------------------
	// complexity: O(n^2)
	// space: O(1)
	total := 0
	for i, val := range height {
		lMax := max(height[:i])
		rMax := max(height[i+1:])
		capacity := min(lMax, rMax) - val
		if capacity > 0 {
			total += capacity
		}
	}
	return total
}

func max(arr []int) int {
	n := 0
	for _, val := range arr {
		if val > n {
			n = val
		}
	}
	return n
}
