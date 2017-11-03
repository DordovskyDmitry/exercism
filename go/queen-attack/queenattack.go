package queenattack

import "strconv"
import "errors"

const testVersion = 2

type coords struct {
	x, y int
}

var letterToInt = map[byte]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
}

func CanQueenAttack(w string, b string) (res bool, e error) {
	defer func() {
		if r := recover(); r != nil {
			res = false
			e = errors.New("Some error")
		}
	}()
	if w == b {
		panic(errors.New("Same place"))
	}
	wCoord := buildCoords(w)
	bCoord := buildCoords(b)
	res = (wCoord.x == bCoord.x) || (wCoord.y == bCoord.y) || (wCoord.x-bCoord.x == wCoord.y-bCoord.y) || (wCoord.x-bCoord.x == bCoord.y-wCoord.y)
	return
}

func buildCoords(w string) coords {
	x := letterToInt[w[0]]
	y, _ := strconv.Atoi(w[1:])
	if y > 8 || x > 8 || y < 1 || x < 1 {
		panic(errors.New("Bad input"))
	}
	return coords{x, y}
}
