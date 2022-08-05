package leetcode

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xs := strconv.Itoa(x)
	n := len(xs)
	var i, j int
	if n%2 == 0 {
		i = n/2 - 1
		j = n / 2
	} else {
		i = n/2 - 1
		j = n/2 + 1
	}

	for ; i >= 0 && j < n; i, j = i-1, j+1 {
		if xs[i] != xs[j] {
			return false
		}
	}
	return true
}
