package main

func inorderTraversal(root *TreeNode) []int {
	leetcode_94_ans = nil
	inorderTraversalRecursiveFunction(root)
	return leetcode_94_ans
}

func inorderTraversalRecursiveFunction(node *TreeNode) {
	if node == nil {
		return
	}
	inorderTraversalRecursiveFunction(node.Left)
	leetcode_94_ans = append(leetcode_94_ans, node.Val)
	inorderTraversalRecursiveFunction(node.Right)
}

var leetcode_94_ans []int
