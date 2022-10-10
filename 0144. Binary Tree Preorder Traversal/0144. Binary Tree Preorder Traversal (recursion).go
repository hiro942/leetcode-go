package leetcode

func preorderTraversalIteration(root *TreeNode) []int {
	var res []int
	stack := make([]*TreeNode, 100)
	top := -1
	if root != nil {
		top++
		stack[top] = root
	}
	for top >= 0 {
		t := stack[top]
		top--
		res = append(res, t.Val)
		if t.Right != nil {
			top++
			stack[top] = t.Right
		}
		if t.Left != nil {
			top++
			stack[top] = t.Left
		}
	}
	return res
}
