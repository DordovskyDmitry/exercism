package accumulate

const testVersion = 1

func Accumulate(a []string, f func(string) string) (r []string) {
	for _, e := range a {
		r = append(r, f(e))
	}
	return r
}
