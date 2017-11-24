package transpose

func Transpose(input []string) []string {
	runeMatrix, rows, columns := initRuneMatrix(input)
	result := make([]string, rows)
	for i := 0; i < rows; i++ {
		flag := false
		bound := columns
		for j := columns - 1; j >= 0; j-- {
			if !flag && runeMatrix[i][j] != '\x00' {
				flag = true
				bound = j + 1
			}
			if flag && runeMatrix[i][j] == '\x00' {
				runeMatrix[i][j] = ' '
			}

		}
		result[i] = string(runeMatrix[i][0:bound])
	}
	return result
}

func initRuneMatrix(input []string) ([][]rune, int, int) {
	rows := 0
	columns := len(input)
	for i := 0; i < columns; i++ {
		if rows < len(input[i]) {
			rows = len(input[i])
		}
	}
	runeMatrix := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		runeMatrix[i] = make([]rune, columns)
	}

	for i, word := range input {
		for j, l := range word {
			runeMatrix[j][i] = l
		}
	}
	return runeMatrix, rows, columns
}
