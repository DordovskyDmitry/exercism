package robotname

import (
	"fmt"
	"math/rand"
)

const testVersion = 1

var maxValue = 26 * 26 * 999

var NamesPool = make(map[int]bool, maxValue)

type Robot struct {
	name string
	num  int
}

func (r *Robot) Name() string {
	if r.name == "" {
		num := generate()
		r.num = num
		r.name = num2name(num)
	}
	return r.name
}

func (r *Robot) Reset() {
	num := generate()
	r.num = num
	r.name = num2name(num)
}

func generate() int {
	num := rand.Intn(maxValue)
	for NamesPool[num] {
		num = rand.Intn(maxValue)
	}
	NamesPool[num] = true
	return num
}

func num2name(num int) string {
	nums := num % 999
	chars := num / 999
	char1 := chars / 26
	char2 := chars % 26
	return fmt.Sprintf("%c%c%03d", char1+'A', char2+'A', nums)
}
