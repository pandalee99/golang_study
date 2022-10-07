package main

func levelOrder(root *TreeNode) [][]int {
	//初始化
	leetcode_102_ans = make([][]int, 0)
	leetcode_102_collection = nil
	leetcode_102_current = 1
	leetcode_102_count = 1
	leetcode_102_future = 0
	leetcode_102_location = 0
	leetcode_102_collection = append(leetcode_102_collection, root)
	levelOrderFunction(leetcode_102_collection[leetcode_102_location])

	return leetcode_102_ans
}

func levelOrderFunction(node *TreeNode) {
	if node == nil {
		return
	}
	//该结点有效
	leetcode_102_current--
	//并入该结点的子节点
	if node.Left != nil {
		leetcode_102_future++
		leetcode_102_collection = append(leetcode_102_collection, node.Left)
	}
	if node.Right != nil {
		leetcode_102_future++
		leetcode_102_collection = append(leetcode_102_collection, node.Right)
	}
	//主要处理逻辑，该行都处理完了，应该进行下一行的处理
	if leetcode_102_current == 0 {

		//往ans添加新数组
		var temp []int
		for i := 0; i < leetcode_102_count; i++ {
			temp = append(temp, leetcode_102_collection[i].Val)
		}
		finalAns := make([]int, len(temp))
		copy(finalAns, temp)
		leetcode_102_ans = append(leetcode_102_ans, finalAns)
		temp = nil

		//已经出了队列的，消失
		leetcode_102_collection = leetcode_102_collection[leetcode_102_count:]
		//下一层的计数器更新
		leetcode_102_current = leetcode_102_future
		//赋予下一层的需要处理的结点数目
		leetcode_102_count = leetcode_102_future

		leetcode_102_future = 0
		//重置处理下标
		leetcode_102_location = 0
	} else {
		//该行还有需要处理的值
		leetcode_102_location++
	}
	//如果下一行没有需要处理的，那就彻底结束了
	if leetcode_102_current == 0 {
		return
	}

	levelOrderFunction(leetcode_102_collection[leetcode_102_location])

}

var leetcode_102_ans [][]int

var leetcode_102_collection []*TreeNode

//当前行结点
var leetcode_102_current int

//收集下一行结点
var leetcode_102_future int

//记录好当前行结点的数量
var leetcode_102_count int

//递归传入的下标
var leetcode_102_location int
