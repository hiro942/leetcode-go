package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return dfs(1, n)
}

func dfs(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}

	var res []*TreeNode
	for i := l; i <= r; i++ {
		leftTrees, rightTrees := dfs(l, i-1), dfs(i+1, r)
		for _, leftRoot := range leftTrees {
			for _, rightRoot := range rightTrees {
				res = append(res, &TreeNode{i, leftRoot, rightRoot})
			}
		}
	}

	return res
}
