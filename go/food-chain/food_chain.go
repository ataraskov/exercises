package foodchain

import (
	"fmt"
	"strings"
)

var (
	animal = []string{
		"",
		"fly",
		"spider",
		"bird",
		"cat",
		"dog",
		"goat",
		"cow",
		"horse",
	}

	second = []string{
		"",
		"",
		"It wriggled and jiggled and tickled inside her.",
		"How absurd to swallow a bird!",
		"Imagine that, to swallow a cat!",
		"What a hog, to swallow a dog!",
		"Just opened her throat and swallowed a goat!",
		"I don't know how she swallowed a cow!",
	}
)

func Verse(v int) string {
	first := fmt.Sprintf("I know an old lady who swallowed a %v.", animal[v])
	last := "I don't know why she swallowed the fly. Perhaps she'll die."
	if v == 8 {
		last = "She's dead, of course!"
	}

	res := []string{}
	res = append(res, first)

	if v > 1 && v < 8 {
		res = append(res, second[v])
	}

	for i := v; i > 1 && i < 8; i-- {
		line := fmt.Sprintf("She swallowed the %v to catch the %v", animal[i], animal[i-1])
		if i == 3 {
			line += " that wriggled and jiggled and tickled inside her"
		}
		line += "."
		res = append(res, line)
	}

	res = append(res, last)
	return strings.Join(res, "\n")
}

func Verses(start, end int) string {
	b := strings.Builder{}
	for i := start; i <= end; i++ {
		b.WriteString(Verse(i) + "\n\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

func Song() string {
	return Verses(1, 8)
}
