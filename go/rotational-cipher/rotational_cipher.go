// Package rotationalcipher implements Caesar cipher
package rotationalcipher

// RotationalCipher converts a string rotationl cipher algorithm
func RotationalCipher(plain string, shiftKey int) string {
	res := []rune{}
	for _, r := range plain {
		if r >= 'a' && r <= 'z' {
			r = 'a' + (r-'a'+rune(shiftKey))%26
		} else if r >= 'A' && r <= 'Z' {
			r = 'A' + (r-'A'+rune(shiftKey))%26
		}
		res = append(res, r)
	}
	return string(res)
}
