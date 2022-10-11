package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r + 1) >> 1
	return &TreeNode{nums[mid], buildBST(nums, l, mid-1), buildBST(nums, mid+1, r)}
}
