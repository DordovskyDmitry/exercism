package railfence

func Encode(s string, n int) (result string) {
	var rails = make([][]rune, n)
	var j = 0
	var downDirection = true
	for _, r := range s {
		rails[j] = append(rails[j], r)
		if downDirection {
			j += 1
		} else {
			j -= 1
		}
		if (j == n-1) || (j == 0) {
			downDirection = !downDirection
		}
	}

	for _, runes := range rails {
		result = result + string(runes)
	}

	return
}

func Decode(s string, n int) (result string) {
	var railsCounts = make([]int, n)

	var j = 0
	var downDirection = true
	for i := 0; i < len(s); i++ {
		railsCounts[j]++
		if downDirection {
			j++
		} else {
			j--
		}
		if j == 0 || j == n-1 {
			downDirection = !downDirection
		}
	}
	from := 0
	var rails = make([]string, n)
	for i, count := range railsCounts {
		rails[i] = s[from : from+count]
		from += count
	}

	j = 0
	downDirection = true
	for i := 0; i < len(s); i++ {
		result += string(rails[j][0])
		rails[j] = rails[j][1:]
		if downDirection {
			j++
		} else {
			j--
		}
		if j == 0 || j == n-1 {
			downDirection = !downDirection
		}
	}

	return
}
