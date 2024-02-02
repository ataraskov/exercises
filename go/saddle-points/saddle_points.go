// Package matrix provides a solution for the Saddle Points on Exercism's Go Track.
package matrix

import (
	"strconv"
	"strings"
)

// Matrix represents a two dimensional map of tree hights.
type Matrix [][]int

// Pair holds coordinates of a tree in a Matrix.
type Pair struct {
	r int
	c int
}

// New creates a Matrix from given string.
func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	m := make(Matrix, len(rows))

	for ri, row := range rows {
		cells := strings.Fields(row)
		m[ri] = make([]int, len(cells))
		for ci, cell := range cells {
			val, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}
			m[ri][ci] = val
		}
	}
	return &m, nil
}

// Saddle returns pairs of good candidates for a tree house.
// NOTICE: not performant solution, better to find min/max once
func (m *Matrix) Saddle() []Pair {
	res := []Pair{}
	for ri, row := range *m {
		for ci, v := range row {
			// highest in the row
			rok := true
			for _, i := range row {
				if v < i {
					rok = false
					break
				}
			}
			if !rok {
				continue
			}
			// lowest in the column
			cok := true
			for i := 0; i < len(*m); i++ {
				if v > (*m)[i][ci] {
					cok = false
					break
				}
			}
			if rok && cok {
				res = append(res, Pair{r: ri + 1, c: ci + 1})
			}

		}
	}
	return res
}
