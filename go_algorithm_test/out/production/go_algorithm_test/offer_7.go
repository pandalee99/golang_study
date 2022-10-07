package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	rootValue := preorder[0]
	ans := &TreeNode{rootValue, nil, nil}
	if inorder[i] > 0 {
		ans.Left = buildTree(preorder[1:inorder[i]+1], inorder[:inorder[i]])
	}
	if inorder[i] != len(inorder)-1 {
		ans.Right = buildTree(preorder[inorder[i]+1:], inorder[inorder[i]+1:])
	}
	return ans

}
