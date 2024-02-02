package romannumerals

import (
	"errors"
	"strings"
)

var roman = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	50:   "L",
	90:   "XC",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var arabic = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("input must be <= 3999")
	}

	res := ""
	var err error
	for _, v := range arabic {
		if input >= v {
			res += strings.Repeat(roman[v], input/v)
			input = input % v
		}
	}
	return res, err
}
