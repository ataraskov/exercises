// Package queenattack provides a solution to the Queen Attack on Exercism's Go Track.
package queenattack

import "fmt"

// CanQueenAttack checks if queens can attach each other from given positions
// Errors if
// * positions are identical for both queens
// * position is out of board limits
func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == blackPosition {
		return false, fmt.Errorf("duplicate position")
	}

	w, err := pos2arr(whitePosition)
	if err != nil {
		return false, err
	}

	b, err := pos2arr(blackPosition)
	if err != nil {
		return false, err
	}

	switch {
	case w[0] == b[0] || w[1] == b[1]: // same column/row
		return true, nil
	case w[0]-b[0] == w[1]-b[1]: // diagonal
		return true, nil
	case w[0] == b[1] && w[1] == b[0]: // diagonal
		return true, nil
	}

	return false, nil
}

// pos2arr converts string position to []int position format
func pos2arr(position string) ([2]int, error) {
	pos := [2]int{}
	for i, r := range position {
		switch i {
		case 0:
			pos[0] = int(r-'a') + 1
		case 1:
			pos[1] = int(r-'1') + 1
		default:
			return [2]int{}, fmt.Errorf("bad position")
		}
	}
	if pos[0] > 8 || pos[0] < 1 || pos[1] > 8 || pos[1] < 1 {
		return [2]int{}, fmt.Errorf("bad position")
	}
	return pos, nil
}
