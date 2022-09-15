package leetcode

import "math"

func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int

	hasInserted := false
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			ans = append(ans, interval)
		} else if newInterval[1] < interval[0] {
			if !hasInserted {
				ans = append(ans, newInterval)
				hasInserted = true
			}
			ans = append(ans, interval)
		} else {
			newInterval[0] = int(math.Min(float64(newInterval[0]), float64(interval[0])))
			newInterval[1] = int(math.Max(float64(newInterval[1]), float64(interval[1])))
		}
	}
	if !hasInserted {
		ans = append(ans, newInterval)
	}

	return ans
}
