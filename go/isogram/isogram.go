// Package isogram contains functions to check if a word or phrase is an isogram.
package isogram

import "strings"

// IsIsogram checks if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, l := range strings.ToLower(word) {
		switch {
		case l == ' ' || l == '-':
			continue
		case seen[l]:
			return false
		default:
			seen[l] = true
		}
	}
	return true
}
