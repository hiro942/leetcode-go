package leetcode

func combine(n int, k int) [][]int {
	var res [][]int
	var path []int
	var dfs func(int)
	dfs = func(u int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		if u > n {
			return
		}
		dfs(u + 1)
		path = append(path, u)
		dfs(u + 1)
		path = path[:len(path)-1]
	}
	dfs(1)
	return res
}
