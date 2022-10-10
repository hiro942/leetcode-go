package leetcode

func postorderTraversalIteration(root *TreeNode) []int {
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
		if t.Left != nil {
			top++
			stack[top] = t.Left
		}
		if t.Right != nil {
			top++
			stack[top] = t.Right
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
