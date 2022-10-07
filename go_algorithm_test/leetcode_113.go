package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	//初始化
	leetcode_113_ans = nil
	leetcode_113_temp = nil
	globalTargetSum = targetSum

	pathSumFunction(root, 0, leetcode_113_temp)

	return leetcode_113_ans
}

func pathSumFunction(node *TreeNode, val int, temp []int) {
	//更新当前值
	val += node.Val
	temp = append(temp, node.Val)
	n := len(temp)
	//如果是叶子结点
	if node.Left == nil && node.Right == nil {
		//如果值符合
		if val == globalTargetSum {
			tempAns := make([]int, n)
			copy(tempAns, temp)
			leetcode_113_ans = append(leetcode_113_ans, tempAns)
		}
	}
	if node.Left != nil {
		pathSumFunction(node.Left, val, temp)
	}

	if node.Right != nil {
		pathSumFunction(node.Right, val, temp)
	}
	temp = temp[:n]
}

var globalTargetSum int

var leetcode_113_ans [][]int

var leetcode_113_temp []int
