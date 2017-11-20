package stringset

import "strings"

type Set struct {
	data map[string]bool
}

func New() Set {
	return Set{map[string]bool{}}
}

func NewFromSlice(s []string) Set {
	sl := map[string]bool{}
	for _, e := range s {
		sl[e] = true
	}
	return Set{sl}
}

func (s Set) String() string {
	if len(s.data) == 0 {
		return "{}"
	}
	sl := make([]string, 0)
	for k := range s.data {
		sl = append(sl, "\""+k+"\"")
	}
	return "{" + strings.Join(sl, ", ") + "}"
}

func (s Set) IsEmpty() bool {
	return len(s.data) == 0
}

func (s Set) Has(str string) bool {
	_, ok := s.data[str]
	return ok
}

func Subset(s1, s2 Set) bool {
	if len(s1.data) > len(s2.data) {
		return false
	}
	for k := range s1.data {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1.data {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return (len(s1.data) == len(s2.data)) && Subset(s1, s2)
}

func (s Set) Add(str string) {
	s.data[str] = true
}

func Intersection(s1, s2 Set) Set {
	s := New()
	for k := range s1.data {
		if s2.Has(k) {
			s.Add(k)
		}
	}
	return s
}

func Difference(s1, s2 Set) Set {
	s := New()
	for k := range s1.data {
		if !s2.Has(k) {
			s.Add(k)
		}
	}
	return s
}

func Union(s1, s2 Set) Set {
	s := New()
	for k := range s1.data {
		s.Add(k)
	}
	for k := range s2.data {
		s.Add(k)
	}
	return s
}
