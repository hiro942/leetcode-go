package leetcode

func trap(height []int) int {
	n := len(height)

	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = height[0], height[n-1]

	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
	}

	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		res += min(left[i], right[i]) - height[i]
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}
