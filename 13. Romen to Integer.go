package leetcode

func romanToInt(s string) int {
	m := map[uint8]int{
		'V': 5,
		'I': 1,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := m[s[0]]

	for i := 1; i < len(s); i++ {
		a, b := m[s[i-1]], m[s[i]]
		if a < b {
			ans += b - 2*a
			continue
		}
		ans += b
	}

	return ans
}
