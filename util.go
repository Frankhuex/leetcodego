package leetcode

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
