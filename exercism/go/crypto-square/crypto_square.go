package cryptosquare

import (
	"strings"
)

func dimensions(len int) (c, r int) {
	c, r = 1, 1
	for {
		switch {
		case c >= r && c-r <= 1 && r*c >= len:
			return c, r
		case c > r:
			r += 1
		default:
			c += 1
		}
	}
}

func normilize(s string) []rune {
	res := []rune{}
	for _, c := range strings.ToLower(s) {
		if c >= 'a' && c <= 'z' || c >= '0' && c <= '9' {
			res = append(res, c)
		}
	}
	return res
}

func Encode(pt string) string {
	chars := normilize(pt)
	if len(chars) == 0 {
		return ""
	}

	res := []string{}
	c, r := dimensions(len(chars))
	for i := len(chars); i < c*r; i++ {
		chars = append(chars, ' ')
	}

	for ci := 0; ci < c; ci++ {
		line := []rune{}
		for ri := 0; ri < r; ri++ {
			line = append(line, chars[c*ri+ci])
		}
		res = append(res, string(line))
	}

	return strings.Join(res, " ")
}
