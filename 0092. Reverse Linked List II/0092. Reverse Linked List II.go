package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := &ListNode{-1, head}
	p := dummy
	for i := 0; i < left-1; i++ {
		p = p.Next
	}
	a, b, c := p, p.Next, p.Next.Next
	for i := left + 1; i <= right; i++ {
		next := c.Next
		c.Next = b
		b = c
		c = next
	}
	a.Next.Next = c
	a.Next = b
	return dummy.Next
}
