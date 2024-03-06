package main

// Problem: https://leetcode.com/problems/container-with-most-water/

// test cases
// []      => 0
// [1]     => 0
// [1,1]   => 1
// [1,9,8] => 8
func maxArea(height []int) int {
	// brute force
	// -------------------------------------------------------------------------
	// complexity: O(n^2)
	// space: O(1)
	area := 0
	for i := range height {
		for j := i + 1; j < len(height); j++ {
			a := min(height[i], height[j]) * (j - i)
			area = max(area, a)
		}
	}
	return area
}
