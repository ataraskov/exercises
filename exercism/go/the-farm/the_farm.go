package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()
	switch {
	case amount < 0 && (err == ErrScaleMalfunction || err == nil):
		err = errors.New("negative fodder")
	case amount > 0 && (err == ErrScaleMalfunction):
		err = nil
		amount *= 2
	case cows == 0:
		err = errors.New("division by zero")
	case cows < 0:
		err = &SillyNephewError{cows: cows}
	}

	if err != nil {
		return 0, err
	}

	return amount / float64(cows), nil
}
