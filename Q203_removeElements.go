package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return dummy.Next
}
