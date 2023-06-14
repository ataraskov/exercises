// Package reverse provides a function to reverse a string
package reverse

// Reverse returns a string in reverse
func Reverse(input string) string {
	out := []rune(input)
	for i, o := len(out)-1, 0; i > o; {
		out[o], out[i] = out[i], out[o]
		i--
		o++
	}
	return string(out)
}
