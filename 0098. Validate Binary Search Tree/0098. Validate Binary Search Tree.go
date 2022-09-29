package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var check func(*TreeNode, int, int) bool
	check = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		if root.Val < min || root.Val > max {
			return false
		}
		return check(root.Left, min, root.Val-1) && check(root.Right, root.Val+1, max)
	}
	return check(root, math.MinInt64, math.MaxInt64)
}
