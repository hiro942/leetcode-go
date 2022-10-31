package leetcode

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		board[x][y] = '@'
		dx, dy := [4]int{0, 1, 0, -1}, [4]int{-1, 0, 1, 0}
		for i := 0; i < 4; i++ {
			a, b := x+dx[i], y+dy[i]
			if a >= 0 && a < m && b >= 0 && b < n && board[a][b] == 'O' {
				dfs(a, b)
			}
		}
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(m-1, j)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '@' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
