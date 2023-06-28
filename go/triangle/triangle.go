// Package triangle provides functionality to determine triangle types
package triangle

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides determins if a triangle is equilateral, isosceles, scalene,
//
//	or not a triangle at all
func KindFromSides(a, b, c float64) Kind {
	switch {
	case a <= 0 || b <= 0 || c <= 0:
		return NaT
	case a+b < c || b+c < a || a+c < b:
		return NaT
	case a == b && a == c && b == c:
		return Equ
	case a == c || a == b || b == c:
		return Iso
	case a != b && a != c && b != c:
		return Sca
	}
	panic("unknown kind")
}
