package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	for _, v := range postorderTraversalRecursion(root.Left) {
		res = append(res, v)
	}
	for _, v := range postorderTraversalRecursion(root.Right) {
		res = append(res, v)
	}
	res = append(res, root.Val)
	return res
}
