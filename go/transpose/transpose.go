package transpose

import "strings"

func Transpose(input []string) []string {
	res := []string{}
	for ci, line := range input {
		for ri, r := range line {
			if len(res) <= ri {
				res = append(res, "")
			}
			if len(res[ri]) < ci {
				res[ri] += strings.Repeat(" ", ci-len(res[ri]))
			}
			res[ri] += string(r)
		}
	}
	return res
}
