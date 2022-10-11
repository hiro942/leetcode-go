package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return buildBST(head, nil)
}

func buildBST(start, end *ListNode) *TreeNode {
	if start == end {
		return nil
	}
	// 快慢指针找中点
	slow, fast := start, start
	for fast != end && fast.Next != end {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return &TreeNode{slow.Val, buildBST(start, slow), buildBST(slow.Next, end)}
}
