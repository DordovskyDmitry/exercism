package perfect

import (
	"errors"
	"math"
)

const testVersion = 1

var ErrOnlyPositive = errors.New("Only positive numbers are allowed")

type Classification string

const (
	ClassificationPerfect   = Classification("ClassificationPerfect")
	ClassificationAbundant  = Classification("ClassificationAbundant")
	ClassificationDeficient = Classification("ClassificationDeficient")
)

func Classify(n uint64) (Classification, error) {
	if n < 1 {
		return "", ErrOnlyPositive
	}
	var s uint64 = 1
	for i := uint64(2); i <= uint64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			d := n / i
			if d == i {
				s += i
			} else {
				s += (d + i)
			}
		}
	}
	if s == n && n != 1 {
		return ClassificationPerfect, nil
	} else if s > n {
		return ClassificationAbundant, nil
	} else {
		return ClassificationDeficient, nil
	}

}
