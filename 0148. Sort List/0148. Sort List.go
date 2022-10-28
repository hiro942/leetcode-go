package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}

	for i := 1; i < n; i *= 2 {
		dummy := &ListNode{-1, nil}
		cur := dummy
		for j := 0; j+i < n; j += i + i {
			p, q := head, head
			for k := 0; k < i && q != nil; k++ {
				q = q.Next
			}
			l, r := 0, 0
			for l < i && r < i && p != nil && q != nil {
				if p.Val <= q.Val {
					cur.Next = p
					cur = cur.Next
					p = p.Next
					l++
				} else {
					cur.Next = q
					cur = cur.Next
					q = q.Next
					r++
				}
			}
			for l < i && p != nil {
				cur.Next = p
				cur = cur.Next
				p = p.Next
				l++
			}
			for r < i && q != nil {
				cur.Next = q
				cur = cur.Next
				q = q.Next
				r++
			}
			head = q
		}
		for head != nil {
			cur.Next = head
			cur = cur.Next
			head = head.Next
		}
		cur.Next = nil
		head = dummy.Next
	}
	return head
}
