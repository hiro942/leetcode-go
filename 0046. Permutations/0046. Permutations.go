package leetcode

var (
	ans     [][]int
	visited []bool
)

func permute(nums []int) [][]int {
	ans = [][]int{}
	visited = make([]bool, len(nums))
	dfs(nums, []int{})
	return ans
}

func dfs(nums []int, path []int) {
	if len(path) == len(nums) {
		ans = append(ans, path)
		return
	}

	for i, val := range nums {
		if !visited[i] {
			visited[i] = true
			dfs(nums, append(path, val))
			visited[i] = false
		}
	}
}
