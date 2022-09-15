package leetcode

func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r) >> 1
		if mid <= x/mid {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
