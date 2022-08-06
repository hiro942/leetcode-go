package leetcode

func intToRoman(num int) (ans string) {
	// 枚举模拟
	units := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(units); i++ {
		for num >= units[i] {
			ans += romans[i]
			num -= units[i]
		}
	}

	return ans
}
