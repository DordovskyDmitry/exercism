package strain

type Ints []int
type Lists [][]int
type Strings []string

func (input Ints) Keep(f func(int) bool) (res Ints) {
	if input == nil {
		return res
	}
	res = make([]int, 0, len(input))
	for _, e := range input {
		if f(e) {
			res = append(res, e)
		}
	}
	return
}

func (input Ints) Discard(f func(int) bool) (res Ints) {
	if input == nil {
		return res
	}
	res = make([]int, 0, len(input))
	for _, e := range input {
		if !f(e) {
			res = append(res, e)
		}
	}
	return
}

func (input Lists) Keep(f func([]int) bool) (res Lists) {
	for _, e := range input {
		if f(e) {
			res = append(res, e)
		}
	}
	return
}

func (input Strings) Keep(f func(string) bool) (res Strings) {
	for _, e := range input {
		if f(e) {
			res = append(res, e)
		}
	}
	return
}
