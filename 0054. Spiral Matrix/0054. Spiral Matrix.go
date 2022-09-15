package leetcode

func spiralOrder(matrix [][]int) []int {
	var ans []int
	n, m := len(matrix), len(matrix[0])
	totalElem := n * m

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	dx, dy := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
	for x, y, to := 0, 0, 0; len(ans) < totalElem; {
		ans = append(ans, matrix[x][y])
		visited[x][y] = true

		a, b := x+dx[to], y+dy[to]
		if a < 0 || a >= n || b < 0 || b >= m || visited[a][b] {
			to = (to + 1) % 4
		}
		x, y = x+dx[to], y+dy[to]
	}

	return ans
}
