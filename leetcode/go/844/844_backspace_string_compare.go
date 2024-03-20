package main

// Problem: https://leetcode.com/problems/backspace-string-compare

func backspaceCompare(s string, t string) bool {
	// brute force
	// -------------------------------------------------------------------------
	// complexity: O(n)
	// space: O(n)
	/*
		buildString := func(s string) string {
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
		return buildString(s) == buildString(t)
	*/

	// two pointers
	// -------------------------------------------------------------------------
	// complexity: O(n)
	// space: O(1)
	skip := func(s string, i int) int {
		skips := 0
		for i >= 0 {
			if s[i] == '#' {
				skips++
				i--
			} else if skips > 0 {
				skips--
				i--
			} else {
				return i
			}
		}
		return i
	}
	si, ti := len(s)-1, len(t)-1
	for si >= 0 || ti >= 0 {
		si = skip(s, si)
		ti = skip(t, ti)
		if si >= 0 && ti >= 0 && s[si] == t[ti] {
			si--
			ti--
			continue
		} else {
			break
		}
	}

	return si == -1 && ti == -1
}
