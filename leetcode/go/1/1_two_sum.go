package main

// Problem: https://leetcode.com/problems/two-sum

// test cases:
// [0,0,1,1], 2 => 2,3
// [1,1], 2     => 0,1
// [1,0,0,1], 2 => 0,3
func twoSum(nums []int, target int) []int {
	// brute force
	// -------------------------------------------------------------------------
	// complexity: O(n^2)
	// space: O(1)
	/*
		for i := range nums {
			for j := i + 1; j < len(nums); j++ {
				if nums[i]+nums[j] == target {
					return []int{i, j}
				}
			}
		}
	*/

	// using map
	// -------------------------------------------------------------------------
	// complexity: O(n)
	// space: O(n)
	m := make(map[int]int)
	for idx, num := range nums {
		if pos, ok := m[target-num]; ok {
			return []int{pos, idx}
		}
		m[num] = idx
	}
	return []int{}
}
