// Package darts implements a solution to the Exercism darts exercise.
package darts

import "math"

const (
	outer  = 10
	middle = 5
	inner  = 1
)

// Score calculates the score of a dart throw.
func Score(x, y float64) int {
	radius := math.Hypot(x, y)
	switch {
	case radius <= inner:
		return 10
	case radius <= middle:
		return 5
	case radius <= outer:
		return 1
	default:
		return 0
	}
}
