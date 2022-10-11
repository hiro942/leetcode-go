package leetcode

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(path string, leftNum, rightNum int)
	dfs = func(path string, leftNum, rightNum int) {
		if rightNum == n {
			res = append(res, path)
			return
		}
		if leftNum < n {
			dfs(path+"(", leftNum+1, rightNum)
		}
		if leftNum > rightNum {
			dfs(path+")", leftNum, rightNum+1)
		}
	}
	dfs("", 0, 0)
	return res
}
