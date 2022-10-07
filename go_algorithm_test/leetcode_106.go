package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	//从后往前扫描
	i := len(inorder) - 1
	//表示最后一个结点
	n := len(postorder) - 1
	//找到根结点
	for ; i >= 0; i-- {
		if inorder[i] == postorder[n] {
			break
		}
	}
	val := postorder[n]
	ans := &TreeNode{val, nil, nil}
	if i > 0 {
		ans.Left = buildTree(inorder[:i], postorder[:i])
	}
	if i != len(inorder)-1 {
		ans.Right = buildTree(inorder[i+1:], postorder[i:n])
	}

	return ans
}
