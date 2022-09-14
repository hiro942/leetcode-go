package leetcode

func addBinary(a string, b string) string {
	var ans string

	if len(a) < len(b) {
		a, b = b, a
	}

	var carry uint8 = 0
	for i, j := len(a)-1, len(b)-1; i >= 0; i, j = i-1, j-1 {
		carry += a[i] - '0'
		if j >= 0 {
			carry += b[j] - '0'
		}
		ans = string(carry%2+'0') + ans
		carry /= 2
	}

	if carry > 0 {
		ans = "1" + ans
	}

	return ans
}
