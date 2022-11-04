package leetcode

func search(nums []int, target int) bool {
	i := len(nums)
	for i >= 0 && nums[i] == nums[0] {
		i--
	}
	if i < 0 {
		return nums[0] == target
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

	if target >= nums[0] {
		l = 0
	} else {
		r = i
	}
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[r] == target
}
