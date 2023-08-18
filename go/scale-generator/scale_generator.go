// package scale provides a solution to Scale Generator exercise on Exercism's Go Track
package scale

import (
	"strings"
)

var (
	SharpScale = []string{
		"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#",
	}
	FlatScale = []string{
		"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab",
	}
	FlatKeys = []string{
		"F", "Bb", "Eb", "Ab", "Db", "Gb",
	}
	FlatKeysMinor = []string{
		"d", "g", "c", "f", "bb", "eb",
	}
)

// index return index if an element in the slice, or -1 if not found
func index(s []string, elem string) int {
	for i, e := range s {
		if strings.EqualFold(elem, e) {
			return i
		}
	}
	return -1
}

func Scale(tonic, interval string) []string {
	var scale []string
	switch {
	case tonic != strings.ToLower(tonic) && index(FlatKeys, tonic) >= 0:
		scale = FlatScale
	case tonic == strings.ToLower(tonic) && index(FlatKeysMinor, tonic) >= 0:
		scale = FlatScale
	default:
		scale = SharpScale
	}

	res := []string{}
	offset := index(scale, tonic)
	if len(interval) == 0 {
		interval = strings.Repeat("m", 11)
	}
	interval = "m" + interval
	for i := 0; i < len(interval); i++ {
		switch interval[i] {
		case 'M':
			offset += 1
		case 'A':
			offset += 2
		}
		index := (offset + i) % len(scale)
		res = append(res, scale[index])
	}

	return res
}
