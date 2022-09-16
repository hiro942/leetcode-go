package leetcode

func lengthOfLastWord(s string) int {
	res := 0
	i := len(s) - 1
	for s[i] == ' ' {
		i--
	}
	for i >= 0 {
		if s[i] == ' ' {
			break
		}
		res++
	}
	return res
}
