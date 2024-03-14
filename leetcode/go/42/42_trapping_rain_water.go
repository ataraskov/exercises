package main

// Problem: https://leetcode.com/problems/trapping-rain-water/

func trap(height []int) int {
	// brute force
	// -------------------------------------------------------------------------
	// complexity: O(n^2)
	// space: O(1)
	/*
		total := 0
		for i, val := range height {
			lMax := max(height[:i])
			rMax := max(height[i+1:])
			capacity := min(lMax, rMax) - val
			if capacity > 0 {
				total += capacity
			}
		}
	*/

	// two pointers
	// -------------------------------------------------------------------------
	// complexity: O(n)
	// space: O(1)
	total := 0
	l, r := 0, len(height)-1
	lMax, rMax := 0, 0

	for l < r {
		lHeight, rHeight := height[l], height[r]
		if lHeight < rHeight {
			if lHeight > lMax {
				lMax = lHeight
			} else {
				total += lMax - lHeight
			}
			l++
		} else {
			if rHeight > rMax {
				rMax = rHeight
			} else {
				total += rMax - rHeight
			}
			r--
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
