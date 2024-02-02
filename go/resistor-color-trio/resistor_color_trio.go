// package resistorcolortrio provides solution for Resistor Color Trio on Exercism's Go Track
package resistorcolortrio

import (
	"fmt"
	"strings"
)

var c2r = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

var n2p = map[int]string{
	0: "",
	1: "kilo",
	2: "mega",
	3: "giga",
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	resistance := 0
	label := ""
	prefix := ""
	for i, c := range colors {
		n := c2r[c]
		switch i {
		case 0:
			resistance += n * 10
		case 1:
			resistance += n
		case 2:
			if resistance > 0 && resistance%10 == 0 {
				n += 1
				resistance = resistance / 10
			}
			prefix = n2p[n/3]
			label = fmt.Sprintf("%d%s", resistance, strings.Repeat("0", n%3))
		default:
			break
		}
	}
	return fmt.Sprintf("%s %sohms", label, prefix)
}
