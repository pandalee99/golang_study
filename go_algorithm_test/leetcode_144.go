package main

func preorderTraversal(root *TreeNode) []int {
	var ans []int

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node != nil {
			ans = append(ans, node.Val)
		} else {
			return
		}
		helper(node.Left)
		helper(node.Right)
	}
	helper(root)
	return ans
}
