// Package spiralmatrix contains a solution for Spiral Matrix exercise on Exercism's Go Track
package spiralmatrix

// SpiralMatrix generate a matrix of given size
func SpiralMatrix(size int) [][]int {
	// init
	res := make([][]int, size)
	for i := range res {
		res[i] = make([]int, size)
	}

	// directions
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}
	// fill
	di := 0
	r := 0
	c := 0
	for n := 1; n <= size*size; n++ {
		res[r][c] = n

		cr := r + dr[di]
		cc := c + dc[di]

		if 0 <= cr && cr < size && 0 <= cc && cc < size && res[cr][cc] == 0 {
			r = cr
			c = cc
		} else {
			di = (di + 1) % 4
			r += dr[di]
			c += dc[di]
		}
	}
	return res
}
