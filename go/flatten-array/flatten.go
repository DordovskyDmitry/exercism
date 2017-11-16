package flatten

func Flatten(in interface{}) []interface{} {
	res := make([]interface{}, 0)

	switch x := in.(type) {
	case []interface{}:
		for _, elem := range x {
			if elem == nil {
				continue
			}
			res = append(res, Flatten(elem)...)
		}
	default:
		res = append(res, x)
	}
	return res
}
