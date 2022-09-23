package leetcode

func subsets(nums []int) [][]int {
	var res [][]int

	n := len(nums)
	for i := 0; i < (1 << n); i++ {
		var path []int
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				path = append(path, nums[j])
			}
		}
		res = append(res, path)
	}

	return res
}