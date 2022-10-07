package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	list *BSTNode
}

type BSTNode struct {
	val  int
	next *BSTNode
}

func Constructor(root *TreeNode) BSTIterator {
	var head *BSTNode
	var list *BSTNode
	var btsNodeFuc func(node *TreeNode)
	btsNodeFuc = func(node *TreeNode) {
		if node == nil {
			return
		}
		btsNodeFuc(node.Left)
		if list == nil {
			list = &BSTNode{node.Val, nil}
			head = list
		} else {
			list.next = &BSTNode{node.Val, nil}
			list = list.next
		}
		btsNodeFuc(node.Right)
	}
	btsNodeFuc(root)
	return BSTIterator{head}
}

func (this *BSTIterator) Next() int {
	val := this.list.val
	this.list = this.list.next
	return val
}

func (this *BSTIterator) HasNext() bool {
	if this.list == nil {
		return false
	} else {
		return true
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
