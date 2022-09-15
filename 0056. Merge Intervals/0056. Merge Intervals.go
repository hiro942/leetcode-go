package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	var ans [][]int

	// sort the intervals by left point
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	l, r := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		a, b := intervals[i][0], intervals[i][1]
		if a <= r {
			r = max(r, b)
		} else {
			ans = append(ans, []int{l, r})
			l, r = a, b
		}
	}
	ans = append(ans, []int{l, r})

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
