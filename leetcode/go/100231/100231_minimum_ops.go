package main

/*
You are given a 0-indexed integer array nums, and an integer k.
In one operation, you can remove one occurrence of the smallest element of nums.
Return the minimum number of operations needed so that all elements of the array are greater than or equal to k.
*/

func minOperations(nums []int, k int) int {
	n := 0
	for _, num := range nums {
		if num < k {
			n++
		}
	}
	return n
}
