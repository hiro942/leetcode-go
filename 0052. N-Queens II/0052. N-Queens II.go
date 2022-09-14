package leetcode

var res int
var col, dg, udg []bool

func totalNQueens(n int) int {
	col = make([]bool, n)
	dg = make([]bool, n+n)
	udg = make([]bool, n+n)
	res = 0
	dfs(0, 0, n)
	return res
}

func dfs(x, qCnt, n int) {
	if x == n {
		if qCnt == n {
			res++
		}
		return
	}

	for y := 0; y < n; y++ {
		if !col[y] && !dg[x-y+n] && !udg[x+y] {
			col[y], dg[x-y+n], udg[x+y] = true, true, true
			dfs(x+1, qCnt+1, n)
			col[y], dg[x-y+n], udg[x+y] = false, false, false
		}
	}
}
