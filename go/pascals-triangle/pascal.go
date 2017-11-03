package pascal

const testVersion = 1

func Triangle(n int) (pt [][]int) {
	pt = make([][]int, n)
	for i := 0; i < n; i++ {
		pt[i] = make([]int, i+1)
		pt[i][0] = 1
		pt[i][i] = 1
		for j := 1; j < i; j++ {
			pt[i][j] = pt[i-1][j-1] + pt[i-1][j]
		}
	}
	return
}
