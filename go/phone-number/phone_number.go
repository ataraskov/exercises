// Package phoneNumber provides functionality for parsing phone numbers.
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

// Number parses a phone number.
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

// AreaCode returns the area code of a phone number.
func AreaCode(phoneNumber string) (string, error) {
	res, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return res[0:3], nil
}

// Format formats a phone number.
func Format(phoneNumber string) (string, error) {
	res, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	res = fmt.Sprintf("(%s) %s-%s", res[:3], res[3:6], res[6:])
	return res, nil
}
