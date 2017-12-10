package matrix

import (
	"strconv"
	"strings"
)

type Pair struct {
	first, last int
}

type Matrix struct {
	m [][]int
}

func New(in string) (*Matrix, error) {
	rows := strings.Split(in, "\n")
	matrix := new(Matrix)
	matrix.m = make([][]int, len(rows))
	for i, row := range rows {
		values := strings.Split(row, " ")
		matrix.m[i] = make([]int, len(values))
		for j, v := range values {
			r, e := strconv.Atoi(v)
			if e != nil {
				return nil, e
			}
			matrix.m[i][j] = r
		}
	}
	return matrix, nil
}

func (matrix *Matrix) Saddle() []Pair {
	maxIndexes := make([]int, len(matrix.m))
	for i, row := range matrix.m {
		maxIndexes[i] = 0
		max := row[0]
		for j, v := range matrix.m[i] {
			if v > max {
				max = v
				maxIndexes[i] = j
			}
		}
	}

	result := make([]Pair, 0, len(matrix.m))

	for i, ind := range maxIndexes {
		isSaddle := true
		for j := 0; j < len(matrix.m); j++ {
			if i == j {
				continue
			}
			if matrix.m[j][ind] < matrix.m[i][ind] {
				isSaddle = false
				break
			}
		}
		if isSaddle {
			result = append(result, Pair{i, ind})
		}
	}
	return result
}
