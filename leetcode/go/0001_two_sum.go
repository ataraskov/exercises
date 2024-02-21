package main

// Problem: https://leetcode.com/problems/two-sum

// test cases:
// [0,0,1,1], 2 => 2,3
// [1,1], 2     => 0,1
// [1,0,0,1], 2 => 0,3
func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
