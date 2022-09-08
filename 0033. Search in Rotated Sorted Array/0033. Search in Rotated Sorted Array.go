package leetcode

func search(nums []int, target int) int {
	// 旋转后分两段，先二分出分界点
	// 二分性质：左边半段元素 >= nums[0]，右边半段元素 < nums[0]

	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] >= nums[0] {
			l = mid
		} else {
			r = mid - 1
		}
	}

	// 判断 target 在哪一段
	if target >= nums[0] {
		l = 0
	} else {
		l, r = l+1, len(nums)-1
	}

	// 二分查找
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// 如果没有第二段的话 l 会越界，需要用 r 来判断
	if nums[r] == target {
		return r
	}

	return -1
}
