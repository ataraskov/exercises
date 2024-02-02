package lsproduct

import (
	"errors"
	"unicode"
)

var (
	ErrTooFewDigits = errors.New("too few digits")
	ErrNegativeSpan = errors.New("negative span")
	ErrNotADigit    = errors.New("digits input must only contain digits")
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return 0, ErrTooFewDigits
	}
	if span < 1 {
		return 0, ErrNegativeSpan
	}

	var res int64 = 0
	var err error = nil

	for i := range digits {
		if i+span > len(digits) {
			break
		}
		var product int64 = 1
		for j, r := range digits {
			switch {
			case j-i >= span:
				break
			case i > j:
				continue
			case !unicode.IsDigit(r):
				return 0, ErrNotADigit
			default:
				product *= int64(r - '0')
			}

		}
		if product > res {
			res = product
		}

	}
	return res, err
}
