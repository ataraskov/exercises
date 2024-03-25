package main

// Problem: https://leetcode.com/problems/valid-palindrome-ii/

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

func validPalindrome(s string) bool {
	if len(s) < 3 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s[left+1:right+1]) || isPalindrome(s[left:right])
		}
		left++
		right--
	}
	return true
}
