package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int

	var q []*TreeNode
	front, rear := 0, -1
	if root != nil {
		q = append(q, root)
		rear++
	}

	for front <= rear {
		var level []int
		m := rear - front + 1
		for i := 0; i < m; i++ {
			t := q[front]
			front++
			level = append(level, t.Val)
			if t.Left != nil {
				q = append(q, t.Left)
				rear++
			}
			if t.Right != nil {
				q = append(q, t.Right)
				rear++
			}
		}
		res = append(res, level)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
