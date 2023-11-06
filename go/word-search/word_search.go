package wordsearch

import (
	"fmt"
	"strings"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	res := make(map[string][2][2]int)

	for _, word := range words {
		sr, sc := -1, -1
		er, ec := -1, -1

		for i, str := range puzzle {
			// case: left to right
			sc = strings.Index(str, word)
			if sc >= 0 {
				sr = i
				er = i
				ec = sc + len(word) - 1
				break
			}
			// case: right to left
			ec = strings.Index(str, reverse(word))
			if ec >= 0 {
				sr = i
				er = i
				sc = sc + len(word)
				break
			}
			// case: diagonal
			// TODO
		}
		if sc >= 0 {
			res[word] = [2][2]int{{sc, sr}, {ec, er}}
		} else {
			return nil, fmt.Errorf("not found: %s", word)
		}
	}
	return res, nil
}

func reverse(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
