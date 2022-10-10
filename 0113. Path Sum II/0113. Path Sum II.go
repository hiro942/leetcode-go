package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			if root.Val == sum {
				res = append(res, append([]int{}, path...))
			}
		} else {
			if root.Left != nil {
				dfs(root.Left, sum-root.Val)
			}
			if root.Right != nil {
				dfs(root.Right, sum-root.Val)
			}
		}
		path = path[:len(path)-1]
	}
	if root != nil {
		dfs(root, targetSum)
	}
	return res
}
