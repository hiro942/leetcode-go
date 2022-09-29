package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	n, m := len(s1), len(s2)

	s1, s2, s3 = " "+s1, " "+s2, " "+s3
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else {
				if s1[i] == s3[i+j] {
					dp[i][j] = dp[i-1][j]
				}
				if s2[j] == s3[i+j] {
					dp[i][j] = dp[i][j] || dp[i][j-1]
				}
			}
		}
	}

	return dp[n][m]
}
