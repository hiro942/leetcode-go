package leetcode

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])

	zeroFirstRow, zeroFirstCol := false, false

	for j := 0; j < m; j++ {
		if matrix[0][j] == 0 {
			zeroFirstRow = true
		}
	}

	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			zeroFirstCol = true
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if zeroFirstRow {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}

	if zeroFirstCol {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}
}
