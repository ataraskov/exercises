// Package anagram contains solution to Anagram exercise of Exercism's Go Track
package anagram

import (
	"strings"
)

// isAnagram returns true if strings are anagrams to each other (case sensitive)
func isAnagram(a, b string) bool {
	if a == b || len(a) != len(b) {
		return false
	}
	for _, s := range strings.Split(a, "") {
		if strings.Count(a, s) != strings.Count(b, s) {
			return false
		}
	}
	return true
}

// Detect returns anagrams from candidates string slice for a given subject
func Detect(subject string, candidates []string) []string {
	res := []string{}
	for i, candidate := range candidates {
		candidate = strings.ToLower(candidate)
		subject = strings.ToLower(subject)
		if isAnagram(subject, candidate) {
			res = append(res, candidates[i])
		}
	}
	return res
}
