package space

import "math"

type Planet string

var planetCoefficient = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

const yearsInSecond = 1.0 / 31557600

func Age(seconds float64, planet Planet) float64 {
	return round2(seconds * yearsInSecond / planetCoefficient[planet])
}

func round2(f float64) (round float64) {
	number := f * 100
	div := number - float64(int(number))
	if div >= 0.5 {
		round = math.Ceil(number)
	} else {
		round = math.Floor(number)
	}
	round = round / 100
	return

}
