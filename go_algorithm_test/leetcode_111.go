package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var queue []*TreeNode
	queue = append(queue, root)
	node := queue[0]
	deep := 1
	i := 0
	current := 1
	count := 1
	future := 0
	for queue != nil {
		node = queue[i]
		current--
		if node.Left == nil && node.Right == nil {
			return deep
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			future++
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			future++
		}
		if current == 0 {
			queue = queue[count:]
			current = future
			count = future
			deep++
			future = 0
			i = 0
		} else {
			i++
		}
	}

	return deep
}
