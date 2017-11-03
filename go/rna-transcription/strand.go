package strand

const testVersion = 3

var RNAMap = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(input string) string {
	var runes []rune
	for _, r := range input {
		runes = append(runes, RNAMap[r])
	}
	return string(runes)
}
