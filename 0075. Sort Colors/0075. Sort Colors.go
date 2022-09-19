package leetcode

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

func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

// Other Method
//func sortColors(nums []int) {
//	for i, j, k := 0, 0, len(nums)-1; j <= k; {
//		if nums[j] == 0 {
//			nums[i], nums[j] = nums[j], nums[i]
//			i++
//			j++
//		} else if nums[j] == 2 {
//			nums[j], nums[k] = nums[k], nums[j]
//			k--
//		} else {
//			j++
//		}
//	}
//}
