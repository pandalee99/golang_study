package main

func flatten(root *TreeNode) {
	//初始化
	leetcode_114_temp = nil
	flattenFunction(root)
	for i := 0; i < len(leetcode_114_temp); i++ {
		prev, curr := leetcode_114_temp[i-1], leetcode_114_temp[i]
		prev.Left, prev.Right = nil, curr
	}

}

func flattenFunction(node *TreeNode) {
	if node == nil {
		return
	}
	leetcode_114_temp = append(leetcode_114_temp, node)
	flattenFunction(node.Left)
	flattenFunction(node.Right)
}

var leetcode_114_temp []*TreeNode
