package leetcode

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	l, r := 0, n-1
	for l < r {
		// find the left index
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// target does not exist
	if nums[l] != target {
		return []int{-1, -1}
	}

	l1 := l

	// find the right index
	l, r = 0, n-1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return []int{l1, l}
}
