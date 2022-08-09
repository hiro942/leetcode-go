package leetcode

func romanToInt(s string) (ans int) {
	var m = map[byte]int{'V': 5, 'I': 1, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	n := len(s)
	for i := range s {
		val := m[s[i]]
		if i < n-1 && val < m[s[i+1]] {
			ans -= val
		} else {
			ans += val
		}
	}

	return ans
}
