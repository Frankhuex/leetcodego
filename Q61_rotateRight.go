package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	length := 1
	for cur.Next != nil {
		cur = cur.Next
		length++
	}
	cur.Next = head
	pre := cur
	for i := 0; i < length-k%length; i++ {
		pre = head
		head = head.Next
	}
	pre.Next = nil
	return head
}
