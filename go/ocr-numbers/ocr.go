package ocr

import (
	"errors"
	"strings"
)

var dict = map[string]string{
	" _ | ||_|": "0",
	"     |  |": "1",
	" _  _||_ ": "2",
	" _  _| _|": "3",
	"   |_|  |": "4",
	" _ |_  _|": "5",
	" _ |_ |_|": "6",
	" _   |  |": "7",
	" _ |_||_|": "8",
	" _ |_| _|": "9",
}

func recognizeDigit(in string) (string, error) {
	lines := strings.Split(in, "\n")
	if len(lines) != 5 { // Last line with \n at the end
		return "", errors.New("wrong height")
	}

	str := ""
	for i := 1; i < 4; i++ {
		if len(lines[i]) > 3 {
			return "", errors.New("wrong line size")
		}
		str += lines[i]
		for j := 0; j < 3-len(lines[i]); j++ {
			str += " "
		}
	}

	res, ok := dict[str]
	if !ok {
		return "?", nil
	}
	return res, nil
}

func Recognize(in string) []string {
	lines := strings.Split(in, "\n")
	digitLines := len(lines) / 4

	res := make([]string, digitLines)
	for i := 0; i < digitLines*4; i += 4 {
		lineLength := len(lines[i+3]) / 3 // all digits have 3 symbols in this line

		for j := 0; j < lineLength*3; j += 3 {
			f := readThree(lines[i+1], j)
			s := readThree(lines[i+2], j)
			t := readThree(lines[i+3], j)

			digit, ok := recognizeDigit("\n" + f + "\n" + s + "\n" + t + "\n")
			if ok != nil {
				panic(ok)
			}
			res[i/4] += digit
		}
	}

	return res
}

func readThree(s string, from int) string {
	to := from + 3
	if to > len(s) {
		to = len(s)
	}
	return s[from:to]
}
