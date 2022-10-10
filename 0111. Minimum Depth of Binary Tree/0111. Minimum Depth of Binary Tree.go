package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := math.MaxInt32
	if root.Right != nil {
		res = min(res, minDepth(root.Right)+1)
	}
	if root.Left != nil {
		res = min(res, minDepth(root.Left)+1)
	}
	if res == math.MaxInt32 {
		// 该节点为叶子节点
		res = 1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	if root.Left == nil && root.Right != nil {
//		return minDepth(root.Right) + 1
//	}
//	if root.Left != nil && root.Right == nil {
//		return minDepth(root.Left) + 1
//	}
//	// 该节点有左右子树 或 为叶子节点
//	return min(minDepth(root.Right), minDepth(root.Left)) + 1
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
