package leetcode

import (
	"sort"
)

var (
	ans     [][]int
	visited []bool
)

func permuteUnique(nums []int) [][]int {
	ans = [][]int{}
	visited = make([]bool, len(nums))
	sort.Ints(nums)
	dfs(nums, []int{})
	return ans
}

func dfs(nums []int, path []int) {
	if len(path) == len(nums) {
		// 这里 append 的必须是 path 元素的复制
		ans = append(ans, append([]int{}, path...))
		return
	}

	for i, val := range nums {
		if visited[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			// 避免两相等的数，先选后再选前的重复情况（需先对nums排序）
			continue
		}
		visited[i] = true
		dfs(nums, append(path, val))
		visited[i] = false
	}
}
