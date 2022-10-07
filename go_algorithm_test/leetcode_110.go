package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leetcode_110_flag = true
	isBalancedFunction110(root, 0)

	return leetcode_110_flag
}

func isBalancedFunction110(node *TreeNode, deep int) int {

	if leetcode_110_flag == false {
		//如果已经不成立了，那么return什么就已经无所谓了
		return 0
	}
	if node == nil {
		return 0
	}

	left := isBalancedFunction110(node.Left, deep)

	right := isBalancedFunction110(node.Right, deep)

	if left-right > 1 || right-left > 1 {
		leetcode_110_flag = false
	}

	//确定好深度
	if right > left {
		deep = right
	} else {
		deep = left
	}
	return deep + 1

}

var leetcode_110_flag bool
