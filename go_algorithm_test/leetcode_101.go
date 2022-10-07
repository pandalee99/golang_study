package main

func isSymmetric(root *TreeNode) bool {
	leetcode_101_flag = true
	isSymmetricFunction(root.Left, root.Right)
	return leetcode_101_flag
}

func isSymmetricFunction(left, right *TreeNode) {
	if leetcode_101_flag == false {
		return
	}
	if left == nil && right == nil {
		return
	}
	if left == nil || right == nil {
		leetcode_101_flag = false
		return
	}
	isSymmetricFunction(left.Left, right.Right)
	if left.Val != right.Val {
		leetcode_101_flag = false
		return
	}
	isSymmetricFunction(left.Right, right.Left)

}

var leetcode_101_flag bool

/*
[1,2,2,2,null,2]
表明，不能通过值的方式去改变


func isSymmetric(root *TreeNode) bool {
	//如果真的是对称的话，那么中序遍历无论是从左到右还是从右到左都应该是一样的
	//初始化
	flag = true
	collection = make([]int, 0)
	isSymmetricLeft(root)
	isSymmetricRight(root)

	return flag
}

func isSymmetricLeft(node *TreeNode) {
	if node == nil {
		return
	}
	isSymmetricLeft(node.Left)
	collection = append(collection, node.Val)
	isSymmetricLeft(node.Right)

}

func isSymmetricRight(node *TreeNode) {
	if node == nil {
		return
	}
	if flag == false {
		return
	}
	isSymmetricRight(node.Right)
	if node.Val == collection[0] {
		//防止溢出
		if len(collection) != 1 {
			collection = collection[1:]
		} else {
			collection = nil
			return
		}
	} else {
		flag = false
		return
	}
	isSymmetricRight(node.Left)
}

var flag bool
var collection []int
*/
