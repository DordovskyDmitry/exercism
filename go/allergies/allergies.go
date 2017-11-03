package allergies

const testVersion = 1

var allergies = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(n uint) []string {
	res := []string{}
	for v, i := range allergies {
		if n&i != 0 {
			res = append(res, v)
		}
	}
	return res
}

func AllergicTo(n uint, allergic string) bool {
	return n&allergies[allergic] != 0
}
