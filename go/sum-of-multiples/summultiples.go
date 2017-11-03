package summultiples

const testVersion = 2

func SumMultiples(limit int, divs ...int) (acc int) {
	for i := 1; i < limit; i++ {
		flag := false
		for _, v := range divs {
			if i%v == 0 {
				flag = true
				break
			}
		}
		if flag {
			acc += i
		}
	}
	return
}
