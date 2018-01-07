package minesweeper

import (
	"bytes"
	"errors"
)

type Board [][]byte

func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}

func (b Board) Count() error {
	height := len(b)
	width := len(b[0])

	for i, line := range b {
		if len(line) != width {
			return errors.New("invalid size")
		}
		if i == 0 || i == height-1 {
			err := checkHorizontalBorder(line)
			if err != nil {
				return err
			}
			continue
		}
		if b[i][0] != '|' {
			return errors.New("invalid symbol")
		}
		for j := 1; j < width-1; j++ {
			if b[i][j] != ' ' && b[i][j] != '*' {
				return errors.New("invalid symbol")
			}
			c := 0
			if b[i][j] == ' ' {
				if b[i-1][j-1] == '*' {
					c += 1
				}
				if b[i-1][j] == '*' {
					c += 1
				}
				if b[i-1][j+1] == '*' {
					c += 1
				}
				if b[i][j-1] == '*' {
					c += 1
				}
				if b[i][j+1] == '*' {
					c += 1
				}
				if b[i+1][j-1] == '*' {
					c += 1
				}
				if b[i+1][j] == '*' {
					c += 1
				}
				if b[i+1][j+1] == '*' {
					c += 1
				}
				if c != 0 {
					b[i][j] = byte(48 + c)
				}
			}
		}

	}
	return nil
}

func checkHorizontalBorder(line []byte) error {
	length := len(line)
	if line[0] != '+' || line[length-1] != '+' {
		return errors.New("invalid symbol")
	}

	for _, l := range line[1 : length-1] {
		if l != '-' {
			return errors.New("invalid symbol")
		}
	}

	return nil
}
