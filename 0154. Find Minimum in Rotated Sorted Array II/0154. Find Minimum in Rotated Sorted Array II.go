package leetcode

func findMin(nums []int) int {
	i := len(nums) - 1
	for i >= 0 && nums[0] == nums[i] {
		i--
	}
	if i < 0 || nums[0] < nums[i] {
		return nums[0]
	}

	l, r := 0, i
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= nums[0] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[r]
}
