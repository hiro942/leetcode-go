package leetcode

func check(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isSame(a, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a += 'a' - 'A'
	}
	if b >= 'A' && b <= 'Z' {
		b += 'a' - 'A'
	}
	return a == b
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !check(s[i]) {
			i++
		}
		for i < j && !check(s[j]) {
			j--
		}
		if i < j && !isSame(s[i], s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
