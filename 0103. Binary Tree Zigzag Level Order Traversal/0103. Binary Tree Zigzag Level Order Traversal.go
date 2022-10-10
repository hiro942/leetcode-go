package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	var Q []*TreeNode
	front, rear := 0, -1
	if root != nil {
		Q = append(Q, root)
		rear++
	}

	direction := true
	for front <= rear {
		m := rear - front + 1
		level := make([]int, m)

		for i := 0; i < m; i++ {
			t := Q[front]
			front++
			if direction {
				level[i] = t.Val
			} else {
				level[m-1-i] = t.Val
			}
			if t.Left != nil {
				Q = append(Q, t.Left)
				rear++
			}
			if t.Right != nil {
				Q = append(Q, t.Right)
				rear++
			}
		}

		direction = !direction
		res = append(res, level)
	}

	return res
}
