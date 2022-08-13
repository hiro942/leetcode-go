package leetcode

var ans []string
var hash = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ans = []string{}
	dfs(digits, "")
	return ans
}

func dfs(digits string, path string) {
	u := len(path)
	if u == len(digits) {
		ans = append(ans, path)
		return
	}

	str := hash[digits[u]-'0']
	for i := range str {
		dfs(digits, path+string(str[i]))
	}
}
