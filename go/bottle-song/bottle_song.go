package bottlesong

import (
	"fmt"
	"strings"
)

var (
	num2str = map[int]string{
		10: "Ten",
		9:  "Nine",
		8:  "Eight",
		7:  "Seven",
		6:  "Six",
		5:  "Five",
		4:  "Four",
		3:  "Three",
		2:  "Two",
		1:  "One",
		0:  "No",
	}

	line1 = "%s green bottle%s hanging on the wall,"
	line2 = line1
	line3 = "And if one green bottle should accidentally fall,"
	line4 = "There'll be %s green bottle%s hanging on the wall."
)

func Recite(startBottles, takeDown int) []string {
	res := []string{}
	for n := startBottles; n > startBottles-takeDown; n-- {
		if len(res) > 0 {
			res = append(res, "")
		}
		s := num2str[n]
		if n != 1 {
			res = append(res, fmt.Sprintf(line1, s, "s"))
			res = append(res, fmt.Sprintf(line2, s, "s"))
		} else {
			res = append(res, fmt.Sprintf(line1, s, ""))
			res = append(res, fmt.Sprintf(line2, s, ""))
		}
		res = append(res, line3)

		reminder := num2str[n-1]
		if n-1 != 1 {
			res = append(res, fmt.Sprintf(line4, strings.ToLower(reminder), "s"))
		} else {
			res = append(res, fmt.Sprintf(line4, strings.ToLower(reminder), ""))
		}

	}
	return res
}
