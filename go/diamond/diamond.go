package diamond

import "errors"

const testVersion = 1

func Gen(c byte) (s string, e error) {
	if c < 'A' || c > 'Z' {
		return "", errors.New("Invalid input")
	}

	dist := (c - 'A')
	length := dist*2 + 1
	for i := byte(0); i < length; i++ {
		for j := byte(0); j < length; j++ {
			if (j+i == dist) || (j-i == dist) || (i-j == dist) || (j+i == length+dist-1) {
				var d byte
				if i <= length/2 {
					d = 'A' + i
				} else {
					d = 'A' + length - i - 1
				}
				s += string(d)
			} else {
				s += " "
			}
		}
		s += "\n"
	}
	return s, nil
}
