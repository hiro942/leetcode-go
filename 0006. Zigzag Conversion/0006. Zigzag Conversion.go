package leetcode

func convert(s string, numRows int) (ans string) {
	if numRows == 1 {
		return s
	}

	mat := make([][]byte, numRows)
	i, to := 0, 1

	for k := range s {
		mat[i] = append(mat[i], s[k])
		i += to

		if i == numRows {
			i, to = i-2, -to
		}
		if i == -1 {
			i, to = 1, -to
		}
	}

	for _, val := range mat {
		ans += string(val)
	}

	return ans
}
