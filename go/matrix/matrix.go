package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

// New returns a Matrix from a string
func New(s string) (Matrix, error) {
	res := Matrix{}
	for _, line := range strings.Split(s, "\n") {
		row := []int{}
		line = strings.TrimSpace(line)
		for _, str := range strings.Split(line, " ") {
			num, err := strconv.Atoi(str)
			if err != nil {
				return Matrix{}, err
			}
			row = append(row, num)
		}
		if len(res) > 0 {
			if len(res[0]) != len(row) {
				return Matrix{}, fmt.Errorf("uneven rows")
			}
		}
		res = append(res, row)
	}
	return res, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	res := make([][]int, len(m[0]))
	for i, rows := range m {
		for j := range rows {
			if i == 0 {
				res[j] = make([]int, len(m))
			}
			res[j][i] = m[i][j]
		}
	}
	return res
}

func (m Matrix) Rows() [][]int {
	res := [][]int{}
	for _, re := range m {
		row := make([]int, len(re))
		for ci, ce := range re {
			row[ci] = ce
		}
		res = append(res, row)
	}
	return res
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row >= len(m) || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
