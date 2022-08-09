package leetcode

func isPalindrome(x int) bool {
	if x < 0 || x > 0 && x%10 == 0 {
		return false
	}

	y := 0
	for y <= x {
		y = y*10 + x%10
		if x == y || x/10 == y {
			return true
		}
		x /= 10
	}

	return false
}
