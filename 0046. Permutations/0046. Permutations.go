package leetcode

func permute(nums []int) [][]int {
	var (
		ans     [][]int
		path    []int
		visited = make([]bool, len(nums))
	)
	dfs(nums, path, visited, &ans)
	return ans
}

func dfs(nums []int, path []int, visited []bool, ans *[][]int) {
	if len(path) == len(nums) {
		*ans = append(*ans, path)
		return
	}

	for i, num := range nums {
		if !visited[i] {
			visited[i] = true
			dfs(nums, append(path, num), visited, ans)
			visited[i] = false
		}
	}
}
