package leetcode

import "math"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i > 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+grid[i][j])
			}
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+grid[i][j])
			}
		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
