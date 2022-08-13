package leetcode

func lengthOfLongestSubstring(s string) (ans int) {
	hash := map[byte]int{}
	for i, j := 0, 0; j < len(s); j++ {
		hash[s[j]]++
		for cnt, ok := hash[s[j]]; ok && cnt > 1; i++ {
			delete(hash, s[i])
		}
		ans = max(ans, j-i+1)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
