package leetcode

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	slow, fast := dummy, dummy
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head2 := slow.Next
	slow.Next = nil
	newHead := sortList(head)
	newHead2 := sortList(head2)
	prev := dummy
	for newHead != nil && newHead2 != nil {
		if newHead.Val < newHead2.Val {
			prev.Next = newHead
			newHead = newHead.Next
		} else {
			prev.Next = newHead2
			newHead2 = newHead2.Next
		}
		prev = prev.Next
	}
	for newHead != nil {
		prev.Next = newHead
		prev = prev.Next
		newHead = newHead.Next
	}
	for newHead2 != nil {
		prev.Next = newHead2
		prev = prev.Next
		newHead2 = newHead2.Next
	}
	return dummy.Next
}
