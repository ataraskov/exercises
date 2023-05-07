package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		switch {
		case r == '\u2757':
			return "recommendation" // ❗
		case r == '🔍':
			return "search" // 🔍 , \u1F50D gives error
		case r == '\u2600':
			return "weather" // ☀
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var result []byte
	for _, r := range log {
		if r == oldRune {
			result = utf8.AppendRune(result, newRune)
		} else {
			result = utf8.AppendRune(result, r)
		}
	}
	return string(result)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
