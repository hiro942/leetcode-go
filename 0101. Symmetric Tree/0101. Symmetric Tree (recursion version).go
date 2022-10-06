package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 两颗子树 root1, root2 镜像条件
// 1. root1.Val == root2.Val
// 2. root1.Left 和 root2.Right 镜像
// 3. root1.Right 和 root2.Left 镜像

func isSymmetricRecursion(root *TreeNode) bool {
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(root1 *TreeNode, root2 *TreeNode) bool {
		if root1 == nil || root2 == nil {
			return root1 == nil && root2 == nil
		}
		return root1.Val == root2.Val && dfs(root1.Left, root2.Right) && dfs(root1.Right, root2.Left)
	}
	return root == nil || dfs(root.Left, root.Right)
}
