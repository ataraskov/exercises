package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	s = strings.ToLower(s)
	res := []rune{}
	count := 0
	for _, r := range s {
		switch {
		case unicode.IsLetter(r):
			r = 'a' + 'z' - r
		case !unicode.IsDigit(r):
			continue
		}
		if count > 0 && count%5 == 0 {
			res = append(res, ' ')
		}
		res = append(res, r)
		count++
	}

	return string(res)
}
