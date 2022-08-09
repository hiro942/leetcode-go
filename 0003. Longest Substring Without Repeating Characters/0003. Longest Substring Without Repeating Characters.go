package leetcode

import "github.com/hiro942/leetcode/utils"

var max = utils.Max

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
