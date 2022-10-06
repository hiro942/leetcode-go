package leetcode

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	for i := 1; i <= rowIndex+1; i++ {
		// 下面要倒序计算，保证res[j-1]还是上一层元素
		for j := i - 1; j >= 0; j-- {
			if j == 0 || j == i-1 {
				res[j] = 1
			} else {
				res[j] += res[j-1]
			}
		}
	}
	return res
}
