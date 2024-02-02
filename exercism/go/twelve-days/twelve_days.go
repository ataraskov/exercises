// Package twelve provides a solution to the twelve days challenge.
package twelve

import (
	"fmt"
	"strings"
)

// Verse
func Verse(i int) string {
	items := []string{
		"",
		"a Partridge in a Pear Tree.",
		"two Turtle Doves, and",
		"three French Hens,",
		"four Calling Birds,",
		"five Gold Rings,",
		"six Geese-a-Laying,",
		"seven Swans-a-Swimming,",
		"eight Maids-a-Milking,",
		"nine Ladies Dancing,",
		"ten Lords-a-Leaping,",
		"eleven Pipers Piping,",
		"twelve Drummers Drumming,",
	}
	days := []string{
		"",
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}
	gifts := []string{}
	day := days[i]
	verse := "On the %s day of Christmas my true love gave to me: %s"
	for ; i > 0; i-- {
		gifts = append(gifts, items[i])
	}
	return fmt.Sprintf(verse, day, strings.Join(gifts, " "))
}

// Song returns the entire song
func Song() string {
	res := []string{}
	for i := 1; i < 13; i++ {
		res = append(res, Verse(i))
	}
	return strings.Join(res, "\n")
}
