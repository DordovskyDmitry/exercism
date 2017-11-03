package brackets

import "errors"

const testVersion = 5

type Stack struct {
	storage []rune
	counter int
}

func (s *Stack) Push(r rune) {
	s.storage = append(s.storage, r)
	s.counter++
}

func (s *Stack) Pop() (r rune, e error) {
	s.counter--
	if s.counter < 0 {
		return r, errors.New("Empty stack error")
	}
	r = s.storage[s.counter]
	s.storage = s.storage[0:s.counter]
	return r, nil
}

var bracketsMap = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

func Bracket(input string) (bool, error) {
	stack := Stack{}
	for _, r := range input {
		if IsOpenBracket(r) {
			stack.Push(r)
			continue
		}
		if IsClosedBracket(r) {
			s, e := stack.Pop()
			if e != nil || bracketsMap[s] != r {
				return false, nil
			}
		}
	}
	if stack.counter != 0 {
		return false, nil
	}
	return true, nil
}

func IsOpenBracket(r rune) bool {
	return r == '(' || r == '[' || r == '{'
}

func IsClosedBracket(r rune) bool {
	return r == ')' || r == ']' || r == '}'
}
