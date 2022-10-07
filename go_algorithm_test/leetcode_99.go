package main

func recoverTree(root *TreeNode) {
	//初始化
	x, y = 0, 0
	nodeA, nodeB = nil, nil
	begin = false
	recoverTreeLeft(root)
	begin = false
	recoverTreeRight(root)
	nodeA.Val, nodeB.Val = nodeB.Val, nodeA.Val

}

func recoverTreeLeft(node *TreeNode) {
	if node == nil {
		return
	}
	recoverTreeLeft(node.Left)
	if begin == false {
		//初始化点，只执行一次
		begin = true
		x = node.Val
	} else {
		//通常逻辑
		if x < node.Val {
			//正确，并获得该结点的值
			x = node.Val
		} else {
			//错误，记录位置并更新
			x = node.Val
			nodeA = node
		}
	}
	recoverTreeLeft(node.Right)

}

func recoverTreeRight(node *TreeNode) {
	if node == nil {
		return
	}
	recoverTreeRight(node.Right)

	if begin == false {
		//初始化点，只执行一次
		begin = true
		y = node.Val
	} else {
		//通常逻辑
		if y > node.Val {
			//正确，并获得该结点的值
			y = node.Val
		} else {
			//错误，记录位置并更新
			y = node.Val
			nodeB = node
		}
	}

	recoverTreeRight(node.Left)
}

var nodeA *TreeNode
var nodeB *TreeNode

var x int
var y int

var begin bool

