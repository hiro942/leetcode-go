package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var res int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		sum = sum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			res += sum
			return
		}
		if node.Left != nil {
			dfs(node.Left, sum)
		}
		if node.Right != nil {
			dfs(node.Right, sum)
		}
	}
	dfs(root, 0)
	return res
}
