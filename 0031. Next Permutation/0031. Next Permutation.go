package leetcode

func nextPermutation(nums []int) {
	n := len(nums)

	// 找到第一次减小的地方
	pos := -1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			pos = i
			break
		}
	}

	if pos == -1 {
		// 为降序数组，整体反转即可
		reverse(nums)
	} else {
		// 从后往前找到最小的比nums[pos]大的数，二者交换
		for i := n - 1; ; i-- {
			if nums[i] > nums[pos] {
				nums[i], nums[pos] = nums[pos], nums[i]
				break
			}
		}
		// 最后反转pos后面的元素
		reverse(nums[pos+1:])
	}
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
