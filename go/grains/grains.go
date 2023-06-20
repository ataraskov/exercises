// Package grains implements a solution to the exercism grains problem
package grains

import (
	"errors"
	"math"
)

const BoardSize = 64

// Square returns the number of grains on a given square
func Square(number int) (uint64, error) {
	if number < 1 || number > BoardSize {
		return 0, errors.New("out of bounds")
	}

	return uint64(math.Pow(2, float64(number-1))), nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
	res := uint64(0)
	for i := 1; i <= BoardSize; i++ {
		g, _ := Square(i)
		res += g
	}
	return res
}
