// Package brackets provides a solution for the exercise Matching Brackets on Exercism's Go Track.
package brackets

// Bracket returns true if brackets in a given string were correctly closed and nested
// otherwise false is returned
func Bracket(input string) bool {
	queue := []rune{}
	for _, r := range input {
		switch {
		case r == '(':
			queue = append([]rune{')'}, queue...)
		case r == '{':
			queue = append([]rune{'}'}, queue...)
		case r == '[':
			queue = append([]rune{']'}, queue...)
		case r == ')' || r == '}' || r == ']':
			if len(queue) == 0 {
				return false
			}
			if queue[0] != r {
				return false
			}
			queue = queue[1:]

		}
	}
	return len(queue) == 0
}
