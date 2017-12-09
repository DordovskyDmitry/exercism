package school

import "sort"

type Grade struct {
	number int
	names  []string
}

type School struct {
	grades [][]string
}

func New() *School {
	grades := make([][]string, 10)
	return &School{grades: grades}
}

func (s *School) Add(name string, n int) {
	if s.grades[n] != nil {
		s.grades[n] = append(s.grades[n], name)
	} else {
		s.grades[n] = []string{name}
	}
}

func (s *School) Grade(grade int) []string {
	x := sort.StringSlice(s.grades[grade])
	sort.Sort(x)
	return x
}

func (s *School) Enrollment() []Grade {
	res := make([]Grade, 0)
	for k, v := range s.grades {
		if len(v) > 0 {
			res = append(res, Grade{number: k, names: s.Grade(k)})
		}
	}
	return res
}
