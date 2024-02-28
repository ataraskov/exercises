package main

// Problem: https://leetcode.com/problems/binary-search

// test cases
// [1]    , 1 => 0
// [1,2,3], 2 => 1
// [1,2,3], 4 => -1
// [1,2,3], 3 => 2

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		middle := start + (end-start)/2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return -1
}
