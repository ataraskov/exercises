package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

var (
	InvalidNumber     = errors.New("invalid number")
	InvalidArea       = errors.New("invalid area")
	InvalidSubscriber = errors.New("invalid subscriber")
)

func Number(phoneNumber string) (string, error) {
	res := []rune{}
	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			if len(res) == 0 && r == '1' {
				continue
			}
			res = append(res, r)
		}
	}

	switch {
	case len(res) != 10:
		return "", InvalidNumber
	case res[0] < '2':
		return "", InvalidArea
	case res[3] < '2':
		return "", InvalidSubscriber
	}

	return string(res), nil
}

func AreaCode(phoneNumber string) (string, error) {
	res, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return res[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	res, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	res = fmt.Sprintf("(%s) %s-%s", res[:3], res[3:6], res[6:])
	return res, nil
}
