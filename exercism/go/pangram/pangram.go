// pacakge pangram provides a function to check if a string is a pangram.
package pangram

import (
	"strings"
)

// IsPangram returns true if the string is a pangram, false otherwise.
func IsPangram(input string) bool {
	input = strings.ToLower(input)
	for c := 'a'; c <= 'z'; c++ {
		if !strings.ContainsRune(input, c) {
			return false
		}
	}
	return true
}
