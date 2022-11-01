package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n, cur := 0, head

	for cur.Next != nil {
		cur = cur.Next
		n++
	}
	n++
	k %= n
	if k == 0 {
		return head
	}

	cur.Next = head
	cur = head
	for i := 0; i < n-k-1; i++ {
		cur = cur.Next
	}
	head = cur.Next
	cur.Next = nil
	return head
}
