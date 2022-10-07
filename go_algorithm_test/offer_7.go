package main

func offfer_7_buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	rootValue := preorder[0]
	ans := &TreeNode{rootValue, nil, nil}
	if i > 0 {
		ans.Left = buildTree(preorder[1:i+1], inorder[:i])
	}
	if i != len(inorder)-1 {
		ans.Right = buildTree(preorder[i+1:], inorder[i+1:])
	}
	return ans

}
