package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil { // 空树是平衡的
		return true
	}
	if abs(maxDepth(root.Left)-maxDepth(root.Right)) > 1 { // check当前树是否平衡
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right) // 当前树平衡的话，还需满足左右子树也是平衡的
}

// maxDepth 返回树的最大高度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
