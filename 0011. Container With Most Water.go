package leetcode

func maxArea(height []int) (ans int) {
	for i, j := 0, len(height)-1; i < j; {
		if height[i] <= height[j] {
			ans = max(ans, (j-i)*height[i])
			i++
		} else {
			ans = max(ans, (j-i)*height[j])
			j--
		}
	}
	return ans
}
