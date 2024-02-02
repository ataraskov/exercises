// Package complexnumbers provides a library of complex number
package complexnumbers

import "math"

// Number is a complex number
type Number struct {
	real, imaginary float64
}

// Real returns the real part of a complex number
func (n Number) Real() float64 {
	return n.real
}

// Imaginary returns the imaginary part of a complex number
func (n Number) Imaginary() float64 {
	return n.imaginary
}

// Add adds two complex numbers
func (n1 Number) Add(n2 Number) Number {
	return Number{
		real:      n1.real + n2.real,
		imaginary: n1.imaginary + n2.imaginary,
	}
}

// Subtract subtracts two complex numbers
func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		real:      n1.real - n2.real,
		imaginary: n1.imaginary - n2.imaginary,
	}
}

// Multiply multiplies two complex numbers
func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		real:      n1.real*n2.real - n1.imaginary*n2.imaginary,
		imaginary: n1.imaginary*n2.real + n1.real*n2.imaginary,
	}
}

// Times multiplies a complex number by provided factor
func (n Number) Times(factor float64) Number {
	return Number{
		real:      n.real * factor,
		imaginary: n.imaginary * factor,
	}
}

// Divide divides two complex numbers
func (n1 Number) Divide(n2 Number) Number {
	return Number{
		real:      (n1.real*n2.real + n1.imaginary*n2.imaginary) / (n2.real*n2.real + n2.imaginary*n2.imaginary),
		imaginary: (n1.imaginary*n2.real - n1.real*n2.imaginary) / (n2.real*n2.real + n2.imaginary*n2.imaginary),
	}
}

// Conjugate returns the conjugate of a complex number
func (n Number) Conjugate() Number {
	return Number{
		real:      n.real,
		imaginary: -n.imaginary,
	}
}

// Abs returns the absolute value of a complex number
func (n Number) Abs() float64 {
	return math.Sqrt(n.real*n.real + n.imaginary*n.imaginary)
}

// Exp returns the exponential of a complex number
func (n Number) Exp() Number {
	return Number{
		real:      math.Exp(n.real) * math.Cos(n.imaginary),
		imaginary: math.Exp(n.real) * math.Sin(n.imaginary),
	}
}
