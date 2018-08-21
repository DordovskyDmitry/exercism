package listops

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int
type IntList []int

func (il IntList) Foldr(fn binFunc, i int) int {
	il = il.Reverse()
	for _, e := range il {
		i = fn(e, i)
	}
	return i
}

func (il IntList) Foldl(fn binFunc, i int) int {
	for _, e := range il {
		i = fn(i, e)
	}
	return i
}

func (il IntList) Filter(fn predFunc) IntList {
	l := IntList{}
	for _, e := range il {
		if fn(e) {
			l = append(l, e)
		}
	}
	return l
}

func (il IntList) Length() int {
	count := 0
	for _, _ = range il {
		count++
	}
	return count
}

func (il IntList) Map(fn unaryFunc) IntList {
	l := IntList{}
	for _, e := range il {
		l = append(l, fn(e))
	}
	return l
}

func (il IntList) Reverse() IntList {
	l := IntList{}
	for _, e := range il {
		l = append(IntList{e}, l...)
	}
	return l
}

func (il IntList) Append(l IntList) IntList {
	return append(il, l...)
}

func (il IntList) Concat(ls []IntList) IntList {
	for _, l := range ls {
		il = il.Append(l)
	}
	return il
}
