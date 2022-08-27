package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head} // 虚拟头节点

	for p := dummy; p.Next != nil && p.Next.Next != nil; p = p.Next.Next {
		a, b := p.Next, p.Next.Next // 需要交换的两个节点
		p.Next = b
		a.Next = b.Next
		b.Next = a
	}

	return dummy.Next
}
