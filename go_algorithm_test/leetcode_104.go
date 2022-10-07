package main

func maxDepth(root *TreeNode) int {
	//初始化
	leetcode_104_ans = 0
	leetcode_104_collection = nil
	leetcode_104_current = 1
	leetcode_104_count = 1
	leetcode_104_future = 0
	leetcode_104_location = 0
	leetcode_104_collection = append(leetcode_104_collection, root)
	maxDepthFunction(root)

	return leetcode_104_ans
}

func maxDepthFunction(node *TreeNode) {
	if node == nil {
		return
	}
	//该结点有效
	leetcode_104_current--
	//并入该结点的子节点
	if node.Left != nil {
		leetcode_104_future++
		leetcode_104_collection = append(leetcode_104_collection, node.Left)
	}
	if node.Right != nil {
		leetcode_104_future++
		leetcode_104_collection = append(leetcode_104_collection, node.Right)
	}
	//主要处理逻辑，该行都处理完了，应该进行下一行的处理
	if leetcode_104_current == 0 {

		//记录深度
		leetcode_104_ans++
		//已经出了队列的，消失
		leetcode_104_collection = leetcode_104_collection[leetcode_104_count:]
		//下一层的计数器更新
		leetcode_104_current = leetcode_104_future
		//赋予下一层的需要处理的结点数目
		leetcode_104_count = leetcode_104_future

		leetcode_104_future = 0
		//重置处理下标
		leetcode_104_location = 0
	} else {
		//该行还有需要处理的值
		leetcode_104_location++
	}
	//如果下一行没有需要处理的，那就彻底结束了
	if leetcode_104_current == 0 {
		return
	}

	maxDepthFunction(leetcode_104_collection[leetcode_104_location])

}

var leetcode_104_collection []*TreeNode

//当前行结点
var leetcode_104_current int

//收集下一行结点
var leetcode_104_future int

//记录好当前行结点的数量
var leetcode_104_count int

//递归传入的下标
var leetcode_104_location int

var leetcode_104_ans int
