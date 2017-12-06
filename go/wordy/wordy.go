package wordy

import (
	"regexp"
	"strconv"
)

var reCommon *regexp.Regexp
var reEval *regexp.Regexp

func init() {
	reCommon, _ = regexp.Compile(`What is (.+)\?`)
	reEval, _ = regexp.Compile(`(-?\d+) (plus|minus|divided by|multiplied by) (-?\d+)(.*)`)
}

func Answer(question string) (int, bool) {
	result := reCommon.FindStringSubmatch(question)
	if len(result) == 0 {
		return 0, false
	}
	return eval(result[1])
}

func eval(question string) (int, bool) {
	ex := reEval.FindStringSubmatch(question)
	if len(ex) == 0 {
		return 0, false
	}
	result, ok := calc(ex)
	if !ok {
		return 0, false
	}
	if ex[4] == "" {
		return result, true
	}
	return eval(strconv.Itoa(result) + ex[4])
}

func calc(ex []string) (int, bool) {
	n, err1 := strconv.Atoi(ex[1])
	if err1 != nil {
		return 0, false
	}
	op := ex[2]
	m, err2 := strconv.Atoi(ex[3])
	if err2 != nil {
		return 0, false
	}
	var result int
	switch op {
	case "plus":
		result = n + m
	case "minus":
		result = n - m
	case "multiplied by":
		result = n * m
	case "divided by":
		result = n / m
	}
	return result, true
}
