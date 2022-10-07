package main

func postorderTraversal(root *TreeNode) []int {
	var ans []int

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node != nil {
			helper(node.Left)
			helper(node.Right)
			ans = append(ans, node.Val)
		} else {
			return
		}

	}
	helper(root)
	return ans
}
