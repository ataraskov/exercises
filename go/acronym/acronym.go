// Package acronym
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	res := ""
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r != '\'' && !unicode.IsLetter(r)
	})
	for _, w := range words {
		res += string(w[0])
	}
	return strings.ToUpper(res)
}
