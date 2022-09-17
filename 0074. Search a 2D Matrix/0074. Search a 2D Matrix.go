package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])

	l, r := 0, n*m-1
	for l < r {
		mid := (l + r) >> 1
		if matrix[mid/m][mid%m] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if matrix[l/m][l%m] != target {
		return false
	}

	return true
}