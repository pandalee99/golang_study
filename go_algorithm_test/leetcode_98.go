package main

func isValidBST(root *TreeNode) bool {
	//初始化
	leetcode_98_x = 0
	leetcode_98_flag = true
	leetcode_98_BeginPoint = false

	isValidBSTFunction(root)

	return leetcode_98_flag
}

func isValidBSTFunction(node *TreeNode) {

	if node == nil {
		return
	}
	if leetcode_98_flag == false {
		return
	}
	isValidBSTFunction(node.Left)
	if leetcode_98_BeginPoint == false {
		//初始化点，只执行一次
		leetcode_98_BeginPoint = true
		leetcode_98_x = node.Val
	} else {
		//通常逻辑
		if leetcode_98_x <= node.Val {
			//正确，并获得该结点的值
			leetcode_98_x = node.Val
		} else {
			//错误，后续也不用进行了
			leetcode_98_flag = false
			return
		}
	}
	isValidBSTFunction(node.Right)

}

var leetcode_98_BeginPoint bool
var leetcode_98_flag bool
var leetcode_98_x int
