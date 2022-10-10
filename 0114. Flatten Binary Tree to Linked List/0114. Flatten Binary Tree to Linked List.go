package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 由于递归先序遍历的时候，回溯到前面节点的时候指针已经被改掉了
// 所以这里从后往前做，以"右左根"的形式遍历，方便一些
func flatten(root *TreeNode) {
	var lastNode *TreeNode = nil
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		dfs(root.Left)
		root.Left = nil
		root.Right = lastNode
		lastNode = root
	}
	dfs(root)
}

//func flatten(root *TreeNode) {
//	var lastNode *TreeNode = nil
//	var dfs func(*TreeNode)
//	dfs = func(root *TreeNode) {
//		if root == nil {
//			return
//		}
//		if lastNode != nil {
//			lastNode.Left = nil
//			lastNode.Right = root
//		}
//		lastNode = root
//		left, right := root.Left, root.Right // 需要事先记录下来
//		dfs(left)
//		dfs(right)
//	}
//	dfs(root)
//}
