package leetcode

// n: [1, 19]
// n <= 18, AC
// n == 19, TLE
func numTreesDFS(n int) int {
	return dfs(1, n)
}

func dfs(l, r int) int {
	if l > r {
		return 1
	}

	num := 0
	for i := l; i <= r; i++ {
		num += dfs(l, i-1) * dfs(i+1, r)
	}

	return num
}
