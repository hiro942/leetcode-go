package leetcode

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// same algorithm, less memory
//func climbStairs(n int) int {
//	a, b := 1, 1
//	for n != 1 {
//		c := a + b
//		a = b
//		b = c
//		n--
//	}
//	return b
//}