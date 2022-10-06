package leetcode

func generate(numRows int) [][]int {
	var res [][]int
	for i := 1; i <= numRows; i++ {
		line := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				line[j] = 1
			} else {
				line[j] = res[i-2][j-1] + res[i-2][j]
			}
		}
		res = append(res, line)
	}
	return res
}