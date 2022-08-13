package leetcode

func longestCommonPrefix(strs []string) string {
	s0, n, flag := strs[0], len(strs), false

	for i := range s0 {
		for j := 1; j < n; j++ {
			if i > len(strs[j])-1 || strs[j][i] != s0[i] {
				flag = true
				break
			}
		}
		if flag {
			return s0[:i]
		}
	}
	return s0
}
