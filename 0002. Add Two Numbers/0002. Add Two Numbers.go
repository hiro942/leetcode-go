package leetcode

import "github.com/hiro942/leetcode/models"

type ListNode = models.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	tail := head
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		tail.Next = &ListNode{sum, nil}
		tail = tail.Next
	}

	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}

	return head.Next
}
