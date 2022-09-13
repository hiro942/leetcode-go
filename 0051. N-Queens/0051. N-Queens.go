package leetcode

var ans [][]string
var solution [][]byte
var col, dg, udg []bool

func solveNQueens(n int) [][]string {
	ans = make([][]string, 0)
	solution = make([][]byte, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			solution[i] = append(solution[i], '.')
		}
	}
	col = make([]bool, n)
	dg = make([]bool, n+n)
	udg = make([]bool, n+n)
	return ans
}

func dfs(x, qCnt, n int) {
	if x == n {
		if qCnt == n {
			tmp := make([]string, n)
			for i := 0; i < n; i++ {
				tmp[i] = string(solution[i])
			}
			ans = append(ans, tmp)
		}
		return
	}

	for y := 0; y < n; y++ {
		if !col[y] && !dg[x-y+n] && !udg[x+y] {
			col[y], dg[x-y+n], udg[x+y] = true, true, true
			solution[x][y] = 'Q'
			dfs(x+1, qCnt+1, n)
			solution[x][y] = '.'
			col[y], dg[x-y+n], udg[x+y] = false, false, false
		}
	}
}
