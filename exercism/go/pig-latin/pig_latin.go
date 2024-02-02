// Package piglatin provides a solution for exercise Pig Latin on Exercism's Go Track
package piglatin

import "strings"

var vowel = []string{"a", "o", "e", "i", "u", "j", "xr", "yt"}

// word returns a word translated to pig-latin
func word(word string) string {
	sound := "ay"
	// starts with vowel
	for _, v := range vowel {
		if strings.HasPrefix(word, v) {
			return word + sound
		}
	}

	// starts with consonant
	prefix := ""
	for i, c := range word {
		for _, v := range vowel {
			if i > 0 && c == 'y' || v == string(c) {
				if strings.HasSuffix(prefix, "q") && c == 'u' {
					prefix += string(c)
				}
				return strings.TrimPrefix(word, prefix) + prefix + sound
			}
		}
		prefix += string(c)
	}

	return word
}

// word returns a sentence translated to pig-latin
func Sentence(sentence string) string {
	words := strings.Split(sentence, " ")
	for i, w := range words {
		words[i] = word(w)
	}
	return strings.Join(words, " ")
}
