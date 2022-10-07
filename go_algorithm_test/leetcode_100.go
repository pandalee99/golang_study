package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	leetcode_100_flag = true
	isSameTreeInorderTraversal(p, q)
	return leetcode_100_flag
}

func isSameTreeInorderTraversal(node1 *TreeNode, node2 *TreeNode) {
	if leetcode_100_flag == false {
		return
	}
	if node1 == nil && node2 == nil {
		return
	}
	if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
		leetcode_100_flag = false
		return
	}
	isSameTreeInorderTraversal(node1.Left, node2.Left)
	if node1.Val != node2.Val {
		leetcode_100_flag = false
		return
	}
	isSameTreeInorderTraversal(node1.Right, node2.Right)

}

var leetcode_100_flag bool
