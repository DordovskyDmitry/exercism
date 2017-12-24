package reverse

func String(s string) string {
	res := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return string(res)
}
