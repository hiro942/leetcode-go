package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	levelStart := root
	for levelStart != nil {
		dummy := new(Node)
		tail := dummy
		for cur := levelStart; cur != nil; cur = cur.Next {
			if cur.Left != nil {
				tail.Next = cur.Left
				tail = tail.Next
			}
			if cur.Right != nil {
				tail.Next = cur.Right
				tail = tail.Next
			}
		}
		levelStart = dummy.Next
	}
	return root
}
