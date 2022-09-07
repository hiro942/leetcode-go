package leetcode

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	n, m := len(num1), len(num2)

	g := make([]int, n+m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			g[i+j] += int((num1[n-1-i] - '0') * (num2[m-1-j] - '0'))
		}
	}

	for i, carry := 0, 0; i < len(g); i++ {
		carry += g[i]
		g[i] = carry % 10
		carry /= 10
	}

	var res string
	for i := len(g) - 1; i >= 0; i-- {
		res += strconv.Itoa(g[i])
	}

	k := 0
	for k < len(res) && res[k] == '0' {
		k++
	}

	if k == len(res) {
		return "0"
	}
	return res[k:]
}
