package leetcode

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	var dfs func(int)

	dfs = func(u int) {
		if u == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		dfs(u + 1) // do not choose nums[u]

		path = append(path, nums[u])
		dfs(u + 1) // choose nums[u]
		path = path[:len(path)-1]
	}
	
	dfs(0)
	return res
}
