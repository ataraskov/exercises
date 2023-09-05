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

	var prev rune
	counter := 1
	encode := func(b bytes.Buffer, c int, r rune) {
		if c > 1 {
			buf.WriteString(fmt.Sprintf("%d", c))
		}
		buf.WriteRune(r)
	}

	for i, r := range input {
		if r == prev {
			counter++
		}
		if i > 0 && r != prev {
			encode(buf, counter, prev)
			counter = 1
		}
		if i+1 == len(input) {
			encode(buf, counter, r)
		}
		prev = r
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
