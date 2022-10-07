package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	//初始化
	leetcode_112_flag = false
	leetcode_112_globalTargetSum = targetSum

	hasPathSumFunction(root, 0)

	return leetcode_112_flag
}

func hasPathSumFunction(node *TreeNode, val int) {
	//如果已经成真了，就不用继续递归了，剪枝
	if leetcode_112_flag == true {
		return
	}
	//更新当前值
	val += node.Val
	//如果是叶子结点
	if node.Left == nil && node.Right == nil {
		//如果值符合
		if val == leetcode_112_globalTargetSum {
			leetcode_112_flag = true
		}
	}
	if node.Left != nil {
		hasPathSumFunction(node.Left, val)
	}
	if node.Right != nil {
		hasPathSumFunction(node.Right, val)
	}

}

var leetcode_112_flag bool
var leetcode_112_globalTargetSum int
