package leetcode

func flatten(root *Node) *Node {
	flatten1(root)
	return root
}

func flatten1(root *Node) (tail *Node) {
	if root == nil {
		return nil
	}
	childTail := flatten1(root.Child)
	nextTail := flatten1(root.Next)
	if root.Next == nil {
		if root.Child == nil {
			return root
		} else {
			root.Next = root.Child
			root.Child.Prev = root
			root.Child = nil
			return childTail
		}
	} else {
		if root.Child == nil {
			return nextTail
		} else {
			next := root.Next
			root.Next = root.Child
			root.Next.Prev = root
			root.Child = nil
			childTail.Next = next
			next.Prev = childTail
			return nextTail
		}
	}
}
