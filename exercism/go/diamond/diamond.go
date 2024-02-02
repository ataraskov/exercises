// Package diamond provides a solution to exercise Diamond on Exercism's Go Track.
package diamond

import (
	"fmt"
	"strings"
)

// Gen returns a diamond shape text starting from A till given letter,
// Where A will be narrowest place (point) of the diamond
// And given letter the widest one
func Gen(char byte) (string, error) {
	switch {
	case char < 'A' || char > 'Z':
		return "", fmt.Errorf("bad char %v proviced", char)
	case char == 'A':
		return "A", nil
	}

	res := []string{}
	// staring from the middle of the diamond
	for c := char; c >= 'A'; c-- {
		filler := " "
		inner := c - 'A'
		outer := char - c
		line := ""

		// left side
		line += strings.Repeat(filler, int(outer))
		line += string(c)
		line += strings.Repeat(filler, int(inner))
		// right side
		// we dont need exact copy, but first char ommited
		if inner > 0 {
			line += strings.Repeat(filler, int(inner-1))
			line += string(c)
		}
		line += strings.Repeat(filler, int(outer))

		res = append(res, line)
		if c < char {
			// add same line as first one if we are not in the middle
			res = append([]string{line}, res...)
		}
	}

	return strings.Join(res, "\n"), nil
}
