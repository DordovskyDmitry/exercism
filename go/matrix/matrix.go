package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	rowCount, colCount int
	data               [][]int
}

func New(in string) (*Matrix, error) {
	rowsData := strings.Split(in, "\n")
	rowCount := len(rowsData)
	m := make([][]int, rowCount)
	colCount := 0
	for i, rowData := range rowsData {
		rowData = strings.TrimSpace(rowData)
		if len(rowData) == 0 {
			return nil, errors.New("Empty line")
		}
		row := strings.Split(rowData, " ")
		length := len(row)
		if colCount == 0 {
			colCount = length
		}
		if colCount != length {
			return nil, errors.New("Not a matrix")
		}
		m[i] = make([]int, colCount)
		for j, e := range row {
			res, err := strconv.Atoi(e)
			if err != nil {
				return nil, err
			}
			m[i][j] = res
		}
	}
	return &Matrix{rowCount, len(m[0]), m}, nil
}

func (m *Matrix) Rows() [][]int {
	copyM := make([][]int, m.rowCount)
	for i, row := range m.data {
		copyM[i] = make([]int, m.colCount)
		for j, v := range row {
			copyM[i][j] = v
		}
	}
	return copyM
}

func (m *Matrix) Cols() [][]int {
	copyM := make([][]int, m.colCount)
	for i := 0; i < m.colCount; i++ {
		copyM[i] = make([]int, m.rowCount)
		for j := 0; j < m.rowCount; j++ {
			copyM[i][j] = m.data[j][i]
		}
	}
	return copyM
}

func (m *Matrix) Set(row, col, val int) bool {
	if row >= 0 && row < m.rowCount && col >= 0 && col < m.colCount {
		m.data[row][col] = val
		return true
	}
	return false
}
