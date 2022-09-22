package leetcode

func isNumber(s string) bool {
	n := len(s)
	hasE, hasPoint := false, false

	for i := 0; i < n; i++ {
		if isDigit(s[i]) {
			continue
		}
		if isE(s[i]) {
			if i == 0 || i == n-1 || hasE {
				return false
			}
			if i == 1 && !isDigit(s[i-1]) {
				return false
			}
			if i > 1 && !isDigit(s[i-1]) && !isDigit(s[i-2]) {
				return false
			}
			hasE = true
			continue
		}
		if isSign(s[i]) {
			if i == n-1 {
				return false
			}
			if i > 0 && !isE(s[i-1]) {
				return false
			}
			continue
		}
		if s[i] == '.' {
			if hasE || hasPoint {
				return false
			}
			if i == 0 && i == n-1 {
				return false
			}
			if i == n-1 && !isDigit(s[i-1]) {
				return false
			}
			if i > 0 && !isSign(s[i-1]) && !isDigit(s[i-1]) {
				return false
			}
			hasPoint = true
			continue
		}
		return false
	}
	return true
}

func isE(c byte) bool {
	if c == 'e' || c == 'E' {
		return true
	}
	return false
}

func isSign(c byte) bool {
	if c == '+' || c == '-' {
		return true
	}
	return false
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
