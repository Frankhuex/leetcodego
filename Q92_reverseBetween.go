package leetcode

import "fmt"

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	prev := dummy
	cur := head
	metLeft := false
	var truePrev *ListNode
	for cur != nil && right > 0 {
		fmt.Println(cur.Val)
		if left == 1 {
			metLeft = true
			truePrev = prev
			fmt.Println("left==1")
		}
		next := cur.Next
		if metLeft {
			cur.Next = prev
			fmt.Print("invert")
			fmt.Print(cur.Val)
		}
		if right == 1 {
			truePrev.Next.Next = next
			truePrev.Next = cur
			fmt.Print("right==1")
			break
		}
		left -= 1
		right -= 1
		prev = cur
		cur = next
	}
	return dummy.Next
}
