package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1r := reverse(l1)
	l2r := reverse(l2)
	head, c := addNodes(l1r, l2r, 0)
	last := head
	for true {
		if l1r != nil {
			l1r = l1r.Next
		}
		if l2r != nil {
			l2r = l2r.Next
		}
		if l1r == nil && l2r == nil && c == 0 {
			return reverse(head)
		}
		last.Next, c = addNodes(l1r, l2r, c)
		last = last.Next
	}
	return nil
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func addNodes(n1, n2 *ListNode, carry int) (*ListNode, int) {
	var v1, v2 int
	if n1 == nil {
		v1 = 0
	} else {
		v1 = n1.Val
	}
	if n2 == nil {
		v2 = 0
	} else {
		v2 = n2.Val
	}
	sum := v1 + v2 + carry
	v := sum % 10
	c := sum / 10
	return &ListNode{
		Val: v,
	}, c
}
