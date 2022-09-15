package leetcode

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	dx, dy := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
	for x, y, to, num := 0, 0, 0, 1; num <= n; num++ {
		ans[x][y] = num

		a, b := x+dx[to], y+dy[to]
		if a < 0 || a >= n || b < 0 || b >= n || ans[a][b] > 0 {
			to = (to + 1) % 4
		}
		x, y = x+dx[to], y+dy[to]
	}

	return ans
}
