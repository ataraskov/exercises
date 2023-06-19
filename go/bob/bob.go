// Package bob is a library for replying to Bob.
package bob

import (
	"strings"
	"unicode"
)

func isShout(phrase string) bool {
	if strings.IndexFunc(phrase, unicode.IsLetter) == -1 {
		return false
	}
	return strings.ToUpper(phrase) == phrase
}

func isQuestion(phrase string) bool {
	return strings.HasSuffix(phrase, "?")
}

func isEmpty(s string) bool {
	return s == ""
}

// Bob replies to what is remarked to Bob.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if isEmpty(remark) {
		return "Fine. Be that way!"
	}

	question := isQuestion(remark)
	shout := isShout(remark)

	switch true {
	case question && shout:
		return "Calm down, I know what I'm doing!"
	case shout:
		return "Whoa, chill out!"
	case question:
		return "Sure."
	default:
		return "Whatever."
	}
}
