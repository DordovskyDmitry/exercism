package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Abbreviate(s string) string {
	if len(s) == 0 {
		return s
	}
	list := regexp.MustCompile("[^a-zA-Z]+").Split(s, -1)
	t := ""
	for _, word := range list {
		t += string([]rune(word)[0])
	}
	return strings.ToUpper(t)
}
