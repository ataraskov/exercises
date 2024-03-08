package main

import "fmt"

// Problem: https://leetcode.com/problems/summary-ranges/

// test cases
// []        => ""
// [1]       => "1"
// [1,2,3]   => "1->3"
// [1,2,5,6] => "1->2,5->6"

func summaryRanges(nums []int) []string {
	if len(nums) < 1 {
		return []string{}
	}

	start, end := nums[0], nums[0]
	result := []string{}

	for _, n := range nums[1:] {
		if n == end+1 {
			end = n
			continue
		}
		result = add(result, start, end)
		start, end = n, n
	}
	return add(result, start, end)
}

func add(result []string, start, end int) []string {
	if start != end {
		return append(result, fmt.Sprintf("%d->%d", start, end))
	}
	return append(result, fmt.Sprintf("%d", start))
}
