package main

/*

 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTreesHelper(1, n)
}

func generateTreesHelper(start int, end int) []*TreeNode {
	if start > end {
		return nil
	}
	var allTreeNode []*TreeNode
	for i := start; i <= end; i++ {
		leftTreeNode := generateTreesHelper(start, i-1)

		rightTreeNode := generateTreesHelper(i+1, end)

		for _, left := range leftTreeNode {
			for _, right := range rightTreeNode {
				currentTreeNode := &TreeNode{i, nil, nil}
				currentTreeNode.Left = left
				currentTreeNode.Right = right
				allTreeNode = append(allTreeNode, currentTreeNode)
			}
		}

	}
	return allTreeNode

}

/*
官方答案

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		leftTrees := helper(start, i-1)
		// 获得所有可行的右子树集合
		rightTrees := helper(i+1, end)
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}


*/
