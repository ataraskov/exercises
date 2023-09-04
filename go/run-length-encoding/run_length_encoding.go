// Package encode provides a solution to Run-Length Encoding on Exercism's Go Track.
package encode

import (
	"bytes"
	"fmt"
	"strconv"
)

// RunLegthEncode returns an encoded string
func RunLengthEncode(input string) string {
	buf := bytes.Buffer{}

	var priv rune
	counter := 1

	for i, r := range input {
		if r == priv {
			counter++
		}
		if i > 0 && r != priv {
			if counter > 1 {
				buf.WriteString(fmt.Sprintf("%d", counter))
			}
			buf.WriteRune(priv)
			counter = 1
		}
		if i+1 == len(input) {
			if counter > 1 {
				buf.WriteString(fmt.Sprintf("%d", counter))
			}
			buf.WriteRune(r)
		}
		priv = r
		/*
			if i == 0 {
				priv = r
				continue
			}
			if i > 0 && r == priv {
				counter++
				if i+1 != len(input) {
					continue
				}
			}
			if counter > 1 {
				buf.WriteString(fmt.Sprintf("%d", counter))
				buf.WriteRune(priv)
				counter = 1
				continue
			}
			buf.WriteRune(r)
			priv = r
		*/
	}
	return buf.String()
}

// RunLengthDecode returns a decoded string
func RunLengthDecode(input string) string {
	buf := bytes.Buffer{}
	digit := false
	repeat := ""

	for _, r := range input {
		if r >= '0' && r <= '9' {
			digit = true
			repeat += string(r)
			continue
		}
		if digit {
			count, err := strconv.Atoi(repeat)
			if err != nil {
				panic("Unable to extract int from " + repeat)
			}
			for j := 0; j < count; j++ {
				buf.WriteRune(r)
			}
			digit = false
			repeat = ""
			continue
		}
		buf.WriteRune(r)
	}
	return buf.String()
}
