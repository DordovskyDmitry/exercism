package say

import "strings"

const testVersion = 1

var basicNumbers = map[uint64]string{
	0:    "zero",
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	1e2:  "hundred",
	1e3:  "thousand",
	1e6:  "million",
	1e9:  "billion",
	1e12: "trillion",
	1e15: "quadrillion",
	1e18: "quintillion",
}

func Say(n uint64) string {
	if n == 0 {
		return basicNumbers[0]
	}
	parts := []string{}

	i := uint64(1)
	for n > 0 {
		rem := n % 1e3
		if rem != 0 {
			if i != 1 {
				parts = append([]string{basicNumbers[i]}, parts...)
			}
			parts = append(SayPart(rem), parts...)
		}
		n = n / 1e3
		i = i * 1e3
	}

	return Join(parts)
}

func SayPart(n uint64) []string {
	n1e2 := n / 100
	rem1e2 := n % 100

	parts := []string{}
	if n1e2 > 0 {
		parts = append(parts, basicNumbers[n1e2], basicNumbers[100])
	}

	if rem1e2 > 0 {
		if rem1e2 < 21 {
			parts = append(parts, basicNumbers[rem1e2])
		} else {
			n1e1 := rem1e2 / 10 * 10
			rem1e1 := rem1e2 % 10
			if rem1e1 > 0 {
				parts = append(parts, Concat(basicNumbers[n1e1], "-", basicNumbers[rem1e1]))
			} else {
				parts = append(parts, basicNumbers[n1e1])
			}
		}

	}
	return parts
}

func Join(parts []string) string {
	return strings.Join(parts, " ")
}

func Concat(strs ...string) string {
	return strings.Join(strs, "")
}

