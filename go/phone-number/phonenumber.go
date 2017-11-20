package phonenumber

import (
	"errors"
	"unicode"
)

func Number(input string) (string, error) {
	str := []rune{}
	for _, v := range input {
		if unicode.IsDigit(v) {
			str = append(str, v)
		}
	}
	length := len(str)
	if length == 10 || (length == 11 && str[0] == '1') {
		res := str[length-10 : length]
		if res[0] == '1' || res[3] == '1' || res[3] == '0' {
			return "", errors.New("Invalid area code")
		}
		return string(res), nil
	}
	return "", errors.New("Invalid input")
}

func AreaCode(input string) (string, error) {
	s, e := Number(input)

	if e != nil {
		return s, e
	}
	return s[0:3], nil
}

func Format(input string) (string, error) {
	s, e := Number(input)

	if e != nil {
		return s, e
	}
	return "(" + s[0:3] + ") " + s[3:6] + "-" + s[6:10], nil
}
