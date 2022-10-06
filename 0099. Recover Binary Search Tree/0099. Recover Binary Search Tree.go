package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var prevNode *TreeNode
	var reversePair []*TreeNode

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if prevNode != nil && prevNode.Val > root.Val {
			reversePair = append(reversePair, prevNode)
			reversePair = append(reversePair, root)
		}
		prevNode = root
		dfs(root.Right)
	}
	dfs(root)

	if len(reversePair) == 2 {
		reversePair[0].Val, reversePair[1].Val = reversePair[1].Val, reversePair[0].Val
	} else {
		reversePair[0].Val, reversePair[3].Val = reversePair[3].Val, reversePair[0].Val
	}
}
