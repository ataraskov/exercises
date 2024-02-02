// Package proverb provides solution to Proverb exercise on Esercism's Go Track.
package proverb

import "fmt"

const (
	stanza = "For want of a %s the %s was lost."
	last   = "And all for the want of a %s."
)

// Proverb ?
func Proverb(rhyme []string) []string {
	res := make([]string, len(rhyme))
	for i := range rhyme {
		if i < len(rhyme)-1 {
			res[i] = fmt.Sprintf(stanza, rhyme[i], rhyme[i+1])
		} else {
			res[i] = fmt.Sprintf(last, rhyme[0])
		}
	}
	return res
}
