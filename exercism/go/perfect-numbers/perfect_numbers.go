// Pacakge perfect provides a solution for the exercise Perfect Numbers on Exercism's Go Track
package perfect

import (
	"errors"
)

type Classification int

const (
	ClassificationDeficient Classification = iota - 1
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("Negative number")

// Classify returns number's classification according to Nicomachus scheme
func Classify(n int64) (Classification, error) {
	if n < 1 {
		return -2, ErrOnlyPositive
	}

	var aliquot_sum int64
	var divisor int64
	for divisor = 1; divisor <= n/2; divisor++ {
		if n%divisor == 0 {
			aliquot_sum += divisor
		}
	}

	switch {
	case n < aliquot_sum:
		return ClassificationAbundant, nil
	case n > aliquot_sum:
		return ClassificationDeficient, nil
	default:
		return ClassificationPerfect, nil
	}
}
