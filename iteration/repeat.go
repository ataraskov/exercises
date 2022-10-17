package iteration

func Repeat(ch string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += ch
	}
	return result
}
