package leetcode

func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	pivot, i, j := nums[l], l-1, r+1
	for i < j {
		for {
			i++
			if nums[i] >= pivot {
				break
			}
		}
		for {
			j--
			if nums[j] <= pivot {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	quickSort(nums, l, j)
	quickSort(nums, j+1, r)
}
