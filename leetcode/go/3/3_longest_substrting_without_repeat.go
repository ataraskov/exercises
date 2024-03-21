package main

// Problem: https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	// brute force
	// -------------------------------------------------------------------------
	// complexity: O(n^2)
	// space: O(n)
	n := 0
	for i := 0; i < len(s); i++ {
		seen := make(map[rune]int)
		for j := i; j < len(s); j++ {
			r := []rune(s)[j]
			if _, found := seen[r]; found {
				break
			}
			if j-i+1 > n {
				n = j - i + 1
			}
			seen[r] = 1
		}
	}
	return n
}
