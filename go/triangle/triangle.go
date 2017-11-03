package triangle

import "math"

const testVersion = 3

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	if NotFinitePositive(a) || NotFinitePositive(b) || NotFinitePositive(c) || BreakTriangleInequalities(a, b, c) {
		return NaT
	}
	if a == b && a == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	} else {
		return Sca
	}
}

func NotFinitePositive(a float64) bool {
	return math.IsNaN(a) || math.IsInf(a, 0) || a <= 0
}

func BreakTriangleInequalities(a, b, c float64) bool {
	return (a+b < c) || (a+c < b) || (b+c < a)
}
