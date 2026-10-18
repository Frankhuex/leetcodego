package leetcode

func rightSideView(root *TreeNode) []int {
	type NodeWithLevel struct {
		N     *TreeNode
		Level int
	}
	queue := []*NodeWithLevel{
		&NodeWithLevel{
			N:     root,
			Level: 1,
		},
	}
	values := []int{}
	level := 0
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top.N == nil {
			continue
		}
		if top.Level > level {
			values = append(values, top.N.Val)
			level++
		}
		queue = append(queue, &NodeWithLevel{
			N:     top.N.Right,
			Level: top.Level + 1,
		}, &NodeWithLevel{
			N:     top.N.Left,
			Level: top.Level + 1,
		})
	}
	return values
}
