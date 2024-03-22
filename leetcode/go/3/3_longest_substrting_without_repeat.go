package main

// Problem: https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	// brute force
	// -------------------------------------------------------------------------
	// complexity: O(n^2)
	// space: O(n)
	/*
		n := 0
		for i := 0; i < len(s); i++ {
			seen := make(map[rune]bool)
			for j := i; j < len(s); j++ {
				r := []rune(s)[j]
				if seen[r] {
					break
				}
				n = max(n, j-i+1)
				seen[r] = true
			}
		}
		return n
	*/

	// sliding window
	// -------------------------------------------------------------------------
	// complexity: O(n)
	// space: O(n)
	n := 0
	seen := make(map[byte]int)
	for left, right := 0, 0; left < len(s) && right < len(s); {
		r := s[right]
		index, found := seen[r]
		if found {
			left = max(left, index+1)
		}
		seen[r] = right
		right++
		n = max(n, right-left)
	}
	return n
}
