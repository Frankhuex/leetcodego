type Node struct {
	Prev *Node
	Next *Node
	Key  int
	Val  int
}

type LRUCache struct {
	Head     *Node
	Tail     *Node
	M        map[int]*Node
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		M:        make(map[int]*Node, capacity),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.M[key]; exists {
		this.MoveToHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.M[key]; exists {
		node.Val = value
		this.MoveToHead(node)
	} else {
		node := &Node{
			Prev: nil,
			Next: this.Head,
			Key:  key,
			Val:  value,
		}
		this.M[key] = node
		if this.Head != nil {
			this.Head.Prev = node
		}
		this.Head = node
		if this.Tail == nil {
			this.Tail = node
		}
		if len(this.M) > this.Capacity {
			delete(this.M, this.Tail.Key)
			this.Tail = this.Tail.Prev
			this.Tail.Next = nil
		}
	}
}

func (this *LRUCache) MoveToHead(node *Node) {
	prev := node.Prev
	next := node.Next
	if node == this.Head {
		return
	}
	node.Next = this.Head
	node.Prev = nil
	if this.Head != nil {
		this.Head.Prev = node
	}
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
	this.Head = node
	if node == this.Tail {
		this.Tail = prev
	}
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */