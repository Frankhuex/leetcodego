package leetcode

func findMode(root *TreeNode) (ans []int) {
	var base, count, mCount int

	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if root.Val == base {
			count++
		} else {
			base, count = root.Val, 1
		}
		if count == mCount {
			ans = append(ans, base)
		} else if count > mCount {
			ans = ans[:0]
			ans = append(ans, base)
			mCount = count
		}
		inorder(root.Right)
	}
	inorder(root)
	return
}
