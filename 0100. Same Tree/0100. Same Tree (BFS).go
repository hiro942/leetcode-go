package leetcode

func isSameTreeBFS(p *TreeNode, q *TreeNode) bool {
	Q := []*TreeNode{p, q}
	front := 0

	for len(Q) > front {
		p, q = Q[front], Q[front+1]
		front += 2

		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil || p.Val != q.Val {
			return false
		}

		Q = append(Q, p.Left)
		Q = append(Q, q.Left)
		Q = append(Q, p.Right)
		Q = append(Q, q.Right)
	}

	return true
}
