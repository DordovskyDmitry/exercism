package pythagorean

import "math"

const testVersion = 1

type Triplet [3]int

func Range(min, max int) (result []Triplet) {
	for m := int(math.Sqrt(float64(max / 2))); m <= int(math.Sqrt(float64(max))); m++ {
		for n := min / (2 * m); n < m; n++ {
			a := m*m + n*n
			b := 2 * m * n
			c := m*m - n*n
			if a <= max && b >= min && c >= min {
				if c > b {
					b, c = c, b
				}
				result = append(result, Triplet{c, b, a})
			}
		}
	}
	return
}

func Sum(p int) (result []Triplet) {
	if p%2 == 1 {
		return
	}
	cache := make(map[Triplet]bool)
	for m := int(math.Sqrt(float64(p / 2))); m > 1; m-- {
		for n := 1; n < m; n++ {
			perimeter := 2 * m * (m + n)
			if p%perimeter == 0 {

				triplet := buildTriplet(m, n, p/perimeter)
				if !cache[triplet] {
					cache[triplet] = true
					result = append(result, triplet)
				}
			}

		}
	}
	return
}

func buildTriplet(m, n, k int) Triplet {
	a := m*m + n*n
	b := 2 * m * n
	c := m*m - n*n

	if c > b {
		b, c = c, b
	}
	return Triplet{k * c, k * b, k * a}
}
