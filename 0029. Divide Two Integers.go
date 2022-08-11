/*
	二分查找
*/

package leetcode

import "math"

func divide(dividend int, divisor int) int {
	// 特判
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		} else {
			return -dividend
		}
	}

	minus := 1
	if dividend < 0 {
		minus *= -1
	}
	if divisor < 0 {
		minus *= -1
	}

	dividendFloat := math.Abs(float64(dividend))
	divisorFloat := math.Abs(float64(divisor))
	l, r := 0.0, dividendFloat/2
	for r-l > 1e-6 {
		mid := (l + r) / 2
		if mid*divisorFloat >= dividendFloat {
			r = mid
		} else {
			l = mid
		}
	}

	// 由于要取整数部分，这里需要取右边界r防止出现答案为2，而 l = 1.99999 r = 2.0000.. 的情况
	return int(math.Floor(r)) * minus
}
