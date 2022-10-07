package main

func levelOrderBottom(root *TreeNode) [][]int {
	//初始化
	leetcode_107_ans = make([][]int, 0)
	leetcode_107_collection = nil
	leetcode_107_current = 1
	leetcode_107_count = 1
	leetcode_107_future = 0
	leetcode_107_location = 0
	leetcode_107_collection = append(leetcode_107_collection, root)
	levelOrderBottomFunction(root)

	//最后交换就好了
	i, j := 0, len(leetcode_107_ans)-1
	for i < j {
		leetcode_107_ans[i], leetcode_107_ans[j] = leetcode_107_ans[j], leetcode_107_ans[i]
		i++
		j--
	}

	return leetcode_107_ans
}

func levelOrderBottomFunction(node *TreeNode) {
	if node == nil {
		return
	}
	//该结点有效
	leetcode_107_current--
	//并入该结点的子节点
	if node.Left != nil {
		leetcode_107_future++
		leetcode_107_collection = append(leetcode_107_collection, node.Left)
	}
	if node.Right != nil {
		leetcode_107_future++
		leetcode_107_collection = append(leetcode_107_collection, node.Right)
	}
	//主要处理逻辑，该行都处理完了，应该进行下一行的处理
	if leetcode_107_current == 0 {

		//往ans添加新数组
		var temp []int
		for i := 0; i < leetcode_107_count; i++ {
			temp = append(temp, leetcode_107_collection[i].Val)
		}
		finalAns := make([]int, len(temp))
		copy(finalAns, temp)
		leetcode_107_ans = append(leetcode_107_ans, finalAns)
		temp = nil

		//已经出了队列的，消失
		leetcode_107_collection = leetcode_107_collection[leetcode_107_count:]
		//下一层的计数器更新
		leetcode_107_current = leetcode_107_future
		//赋予下一层的需要处理的结点数目
		leetcode_107_count = leetcode_107_future

		leetcode_107_future = 0
		//重置处理下标
		leetcode_107_location = 0
	} else {
		//该行还有需要处理的值
		leetcode_107_location++
	}
	//如果下一行没有需要处理的，那就彻底结束了
	if leetcode_107_current == 0 {
		return
	}

	levelOrderBottomFunction(leetcode_107_collection[leetcode_107_location])
}

var leetcode_107_ans [][]int

var leetcode_107_collection []*TreeNode

//当前行结点
var leetcode_107_current int

//收集下一行结点
var leetcode_107_future int

//记录好当前行结点的数量
var leetcode_107_count int

//递归传入的下标
var leetcode_107_location int
