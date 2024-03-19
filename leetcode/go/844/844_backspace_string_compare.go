package main

// Problem: https://leetcode.com/problems/backspace-string-compare

func buildString(s string) string {
	b := []byte{}
	for _, c := range s {
		if c != '#' {
			b = append(b, byte(c))
		} else {
			if len(b) > 0 {
				b = b[:len(b)-1]
			}
		}
	}
	return string(b)
}

func backspaceCompare(s string, t string) bool {
	// brute force
	// -------------------------------------------------------------------------
	// complexity: O(n)
	// space: O(n)
	return buildString(s) == buildString(t)
}
