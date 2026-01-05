package leetcode

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Val: 0, Next: head}
	lastSort := head
	cur := head.Next
	for cur != nil {
		if lastSort.Val <= cur.Val {
			lastSort = cur
			cur = cur.Next
		} else {
			prev := dummy
			for prev.Next.Val < cur.Val {
				prev = prev.Next
			}
			nextN := prev.Next
			lastSort.Next = cur.Next
			prev.Next = cur
			cur.Next = nextN
			cur = lastSort.Next
		}
	}
	return dummy.Next

}
