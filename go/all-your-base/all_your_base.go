package allyourbase

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrInputBase  = errors.New("input base must be >= 2")
	ErrOutputBase = errors.New("output base must be >= 2")
	ErrBadDigit   = errors.New("all digits must satisfy 0 <= d < input base")
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, ErrInputBase
	}
	if outputBase < 2 {
		return []int{}, ErrOutputBase
	}

	n, err := fromDigits(inputBase, inputDigits)
	if err != nil {
		return []int{}, err
	}
	res := toDigits(outputBase, n)
	return res, nil
}

func fromDigits(base int, digits []int) (int, error) {
	res := 0
	for i, d := range digits {
		if d >= base || d < 0 {
			return 0, ErrBadDigit
		}
		power := len(digits) - i - 1
		res += d * int(math.Pow(float64(base), float64(power)))
	}
	return res, nil
}

func toDigits(base int, num int) []int {
	res := []int{}
	if num == 5 && base == 2 {
		fmt.Println("debug")
	}
	for n := num; n >= 0; n = int(math.Floor(float64(n) / float64(base))) {
		if n < base {
			res = append([]int{n}, res...)
			break
		}
		res = append([]int{n % base}, res...)
	}
	return res
}
