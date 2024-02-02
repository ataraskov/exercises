// Pacage wordcount provides a solution for Word Count exercise on Exercism's Go Track
package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

// WordCount returns word frequence map a given phase
func WordCount(phrase string) Frequency {
	res := Frequency{}
	phrase = strings.ToLower(phrase)
	for _, w := range strings.FieldsFunc(phrase, notWordRune) {
		w = strings.Trim(w, "'")
		if len(w) > 0 {
			res[w]++
		}
	}

	return res
}

func notWordRune(r rune) bool {
	return !wordRune(r)
}

func wordRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\''
}
