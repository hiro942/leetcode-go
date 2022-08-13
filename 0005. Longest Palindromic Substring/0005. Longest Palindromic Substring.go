package leetcode

func longestPalindrome(s string) (ans string) {
	for k := range s {
		// 奇数长度回文串
		i, j := k-1, k+1
		for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
		}
		if j-i-1 > len(ans) {
			ans = s[i+1 : j]
		}

		// 偶数长度回文串
		i, j = k, k+1
		for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
		}
		if j-i-1 > len(ans) {
			ans = s[i+1 : j]
		}
	}

	return ans
}
