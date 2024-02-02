// Package isbn provides functions for validating ISBN numbers.
package isbn

import "strings"

// isVadidISBN checks if a string is a valid ISBN number.
func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)

	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, c := range isbn {
		if i == 9 && (c == 'X' || c == 'x') {
			sum += 10
			continue
		}
		sum += int(c-'0') * (10 - i)
	}

	return sum%11 == 0
}
