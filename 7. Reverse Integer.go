package leetcode

import "math"

func reverse(x int) (ans int) {
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		ans = ans*10 + digit
	}
	return
}

//// method 2
//func reverse(x int) int {
//	s := []byte(strconv.Itoa(x))
//	i, j := 0, len(s)-1
//	if s[0] == '-' {
//		i = 1
//	}
//
//	for ; i < j; i, j = i+1, j-1 {
//		tmp := s[i]
//		s[i] = s[j]
//		s[j] = tmp
//	}
//
//	ans, _ := strconv.Atoi(string(s))
//	if math.Abs(float64(ans)) >= 1<<31 {
//		ans = 0
//	}
//
//	return ans
//}
