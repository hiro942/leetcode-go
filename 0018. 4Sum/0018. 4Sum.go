/*
	两层for枚举答案 i、j、k、l 四个数中的 i，j。双指针扫描 k 和 l
	复杂度 O(n^3)  n为nums长度
*/

package leetcode

import "sort"

func fourSum(nums []int, target int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k, l := j+1, n-1; k < l; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				for k < l && nums[i]+nums[j]+nums[k]+nums[l] > target {
					l--
				}
				if k != l && nums[i]+nums[j]+nums[k]+nums[l] == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
				}
			}
		}
	}
	return ans
}
