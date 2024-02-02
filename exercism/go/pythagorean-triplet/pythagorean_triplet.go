package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	res := []Triplet{}

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					res = append(res, Triplet{a, b, c})
					break
				}
				if a*a+b*b < c*c {
					break
				}
			}
		}
	}
	return res
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	res := []Triplet{}
	min := 1
	for a := min; a <= p/2; a++ {
		for b := a; b <= p/2; b++ {
			c := p - a - b
			if a+b+c == p && a*a+b*b == c*c {
				res = append(res, Triplet{a, b, c})
			}
		}
	}
	return res
}
