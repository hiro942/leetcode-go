package leetcode

import "github.com/hiro942/leetcode/utils"

var min = utils.Min

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 0 {
		middleLeft := find(nums1, 0, nums2, 0, n/2)
		middleRight := find(nums1, 0, nums2, 0, n/2+1)
		return (middleLeft + middleRight) / 2
	}
	return find(nums1, 0, nums2, 0, n/2+1)
}

// 再下标i开始的nums1和下标j开始的nums2中找出第k大的
func find(nums1 []int, i int, nums2 []int, j int, k int) float64 {
	// 假设nums1是有效长度更短的那个
	if len(nums1)-i > len(nums2)-j {
		return find(nums2, j, nums1, i, k)
	}

	// 如果 nums1 可找的长度为 0
	if i == len(nums1) {
		return float64(nums2[j+k-1])
	}

	if k == 1 {
		return float64(min(nums1[i], nums2[j]))
	}

	ki := min(len(nums1)-1, i+k/2-1) // nums1中第k/2元素下标
	kj := j + k - k/2 - 1
	if nums1[ki] > nums2[kj] {
		return find(nums1, i, nums2, kj+1, k-(kj-j+1))
	} else {
		return find(nums1, ki+1, nums2, j, k-(ki-i+1))
	}
}
