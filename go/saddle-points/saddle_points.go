// Package matrix provides a solution for the Saddle Points on Exercism's Go Track.
package matrix

import (
	"strconv"
	"strings"
)

// Matrix represents a two dimensional map of tree hights.
type Matrix struct {
	arr [][]int
}

// Pair holds coordinates of a tree in a Matrix.
type Pair struct {
	r int
	c int
}

// New creates a Matrix from given string.
func New(s string) (*Matrix, error) {
	if len(s) == 0 {
		return &Matrix{}, nil
	}
	arr := [][]int{}
	for ri, row := range strings.Split(s, "\n") {
		arr = append(arr, []int{})
		for _, v := range strings.Split(row, " ") {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			arr[ri] = append(arr[ri], n)
		}
	}
	return &Matrix{arr: arr}, nil
}

// Saddle returns pairs of good candidates for a tree house.
func (m *Matrix) Saddle() []Pair {
	res := []Pair{}
	for ri, row := range m.arr {
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
			for i := 0; i < len(m.arr); i++ {
				if v > m.arr[i][ci] {
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
