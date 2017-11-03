package secret

const testVersion = 2

var degreeValues = [4]string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(n uint) (s []string) {
	reverse := false
	if n > 16 {
		reverse = true
	}
	n %= 16
	i := 0
	for n > 0 {
		if n&(1<<1) == 1 {
			if reverse {
				s = append([]string{degreeValues[i]}, s...)
			} else {
				s = append(s, degreeValues[i])
			}
		}
		n /= 2
		i++
	}
	return
}
