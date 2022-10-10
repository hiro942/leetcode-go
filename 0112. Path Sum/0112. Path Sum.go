package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil {
		return targetSum == root.Val
	}
	if root.Right != nil && hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	if root.Left != nil && hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}
	return false
}
