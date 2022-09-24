package leetcode

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	dx, dy := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
	var dfs func(int, int, int) bool
	dfs = func(x, y, u int) bool {
		if board[x][y] != word[u] {
			return false
		}
		if u == len(word)-1 {
			return true
		}
		tmp := board[x][y]
		board[x][y] = '.'
		for i := 0; i < 4; i++ {
			a, b := x+dx[i], y+dy[i]
			if a >= 0 && a < m && b >= 0 && b < n && board[a][b] == word[u+1] {
				if dfs(a, b, u+1) {
					return true
				}
			}
		}
		board[x][y] = tmp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
