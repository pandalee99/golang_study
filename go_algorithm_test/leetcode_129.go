package main

func sumNumbers(root *TreeNode) int {
	leetcode_129_ans = 0
	sumNumbersFunction(root, 0)
	return leetcode_129_ans
}

func sumNumbersFunction(node *TreeNode, val int) {
	if node.Left == nil && node.Right == nil {
		leetcode_129_ans += val*10 + node.Val
		return
	}
	if node.Left != nil {
		sumNumbersFunction(node.Left, val*10+node.Val)
	}
	if node.Right != nil {
		sumNumbersFunction(node.Right, val*10+node.Val)

	}
}

var leetcode_129_ans int
