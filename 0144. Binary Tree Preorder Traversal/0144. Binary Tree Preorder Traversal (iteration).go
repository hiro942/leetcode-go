package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = append(res, root.Val)
	for _, v := range preorderTraversalRecursion(root.Left) {
		res = append(res, v)
	}
	for _, v := range preorderTraversalRecursion(root.Right) {
		res = append(res, v)
	}
	return res
}
