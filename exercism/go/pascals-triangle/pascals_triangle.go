// Package pascal provides a solution to exercise Pascal's Triangle on Exercism's Go Track.
package pascal

// Triangle returns a Pascal's triangle up to a given number of rows.
func Triangle(n int) [][]int {
	t := make([][]int, n)

	for i := 0; i < n; i++ {
		t[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == i || j == 0 {
				t[i][j] = 1
			} else {
				t[i][j] = t[i-1][j-1] + t[i-1][j]
			}
		}
	}
	return t
}
