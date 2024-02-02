// package house contains a solution for House exercise on Exercism's Go Track
package house

import "strings"

var (
	who = []string{
		"house",
		"malt",
		"rat",
		"cat",
		"dog",
		"cow with the crumpled horn",
		"maiden all forlorn",
		"man all tattered and torn",
		"priest all shaven and shorn",
		"rooster that crowed in the morn",
		"farmer sowing his corn",
		"horse and the hound and the horn",
	}
	action = []string{
		"Jack built.",
		"lay in",
		"ate",
		"killed",
		"worried",
		"tossed",
		"milked",
		"kissed",
		"married",
		"woke",
		"kept",
		"belonged to",
	}
)

// Verse returns an Nth verse of the house song
func Verse(v int) string {
	res := []string{}

	for i := v - 1; i >= 0; i-- {
		switch {
		case v == 1:
			res = append(res, "This is the "+who[i]+" that "+action[i])
			continue
		case i == v-1:
			res = append(res, "This is the "+who[i])
		}

		switch {
		case i == 1:
			res = append(res, "that "+action[i]+" the "+who[i-1]+" that "+action[i-1])
		case i > 1:
			res = append(res, "that "+action[i]+" the "+who[i-1])
		}
	}

	return strings.Join(res, "\n")
}

// Song returns the house song
func Song() string {
	res := []string{}
	for i := 0; i < len(who); i++ {
		res = append(res, Verse(i+1))
	}
	return strings.Join(res, "\n\n")
}
