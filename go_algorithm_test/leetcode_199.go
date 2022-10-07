package main

func rightSideView(root *TreeNode) []int {
	var ans []int
	//直接使用层序遍历去解决问题
	var bfs func(node *TreeNode)
	var collection []*TreeNode
	collection = append(collection, root)
	count, next, location := 1, 0, 0
	bfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		count--
		if node.Left != nil {
			collection = append(collection, node.Left)
			next++
		}
		if node.Right != nil {
			collection = append(collection, node.Right)
			next++
		}
		if count == 0 {
			ans = append(ans, node.Val)
			collection = collection[len(collection)-next:]
			count = next
			next = 0
			location = 0
		} else {
			location++
		}
		if count == 0 {
			return
		}
		bfs(collection[location])
	}

	bfs(root)
	return ans
}
