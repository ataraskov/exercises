// Package minesweeper provides a solution for Minesweeper on Exercism's Go Track.
package minesweeper

import (
	"unicode"
)

// Annotate returns an annotated board
func Annotate(board []string) []string {

	cell := make([][]rune, len(board))
	for i := range board {
		cell[i] = ([]rune)(board[i])
	}
	for l := 0; l < len(cell); l++ {
		for c := 0; c < len(cell[l]); c++ {
			if cell[l][c] == '*' {
				markAround(cell, l, c)
			}
		}
	}

	res := []string{}
	for l := range cell {
		res = append(res, (string)(cell[l]))
	}
	return res
}

func markAround(cell [][]rune, l, c int) {
	for i := l - 1; i < l+2; i++ {
		for j := c - 1; j < c+2; j++ {
			if i < 0 || j < 0 || i >= len(cell) || j >= len(cell[i]) {
				continue
			}
			cell[i][j] = inc(cell[i][j])
		}
	}
}

func inc(r rune) rune {
	switch {
	case r == '*':
		return '*'
	case r == ' ':
		return '1'
	case unicode.IsDigit(r):
		return r + 1
	default:
		return ' '
	}
}
