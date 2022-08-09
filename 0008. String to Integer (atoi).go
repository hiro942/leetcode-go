package leetcode

import (
	"math"
)

func myAtoi(s string) (ans int) {
	i := 0
	for ; i < len(s) && s[i] == ' '; i++ {
	}

	if i == len(s) {
		return 0
	}

	minus := 1
	if s[i] == '-' {
		minus = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		ans = ans*10 + int(s[i]-'0')
		if minus == 1 && ans > math.MaxInt32 {
			return math.MaxInt32
		}
		if minus == -1 && -ans < math.MinInt32 {
			return math.MinInt32
		}
	}

	return ans * minus
}
