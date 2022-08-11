package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	bestSum := math.MaxInt32
	n := len(nums)

	updateBestSum := func(sum int) {
		if abs(sum-target) < abs(bestSum-target) {
			bestSum = sum
		}
	}

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, n-1; j < k; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[i]+nums[j]+nums[k] > target {
				k--
			}
			if k+1 < n && j != k+1 {
				updateBestSum(nums[i] + nums[j] + nums[k+1])
			}
			if j != k {
				updateBestSum(nums[i] + nums[j] + nums[k])
			}
		}
	}
	return bestSum
}
