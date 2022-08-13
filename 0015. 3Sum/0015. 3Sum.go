/*
	枚举一层 + 双指针一层   复杂度 O(n^2)  n为nums长度
*/

package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, n-1; j < k; j++ { // 双指针
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}
			// 需要排除 j == k 的元素重复使用情况
			if j != k && nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}
