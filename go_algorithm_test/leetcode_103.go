package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	//初始化
	leetcode_103_ans = make([][]int, 0)
	leetcode_103_collection = nil
	leetcode_103_current = 1
	leetcode_103_count = 1
	leetcode_103_future = 0
	leetcode_103_location = 0
	leetcode_103_mutex = true
	leetcode_103_collection = append(leetcode_103_collection, root)
	zigzagLevelOrderFunction(leetcode_103_collection[leetcode_103_location])

	return leetcode_103_ans
}

func zigzagLevelOrderFunction(node *TreeNode) {
	if node == nil {
		return
	}
	//该结点有效
	leetcode_103_current--
	//并入该结点的子节点
	if node.Left != nil {
		leetcode_103_future++
		leetcode_103_collection = append(leetcode_103_collection, node.Left)
	}
	if node.Right != nil {
		leetcode_103_future++
		leetcode_103_collection = append(leetcode_103_collection, node.Right)
	}
	//主要处理逻辑，该行都处理完了，应该进行下一行的处理
	if leetcode_103_current == 0 {

		//往ans添加新数组
		var temp []int

		//由此产生锯齿形
		if leetcode_103_mutex == true {
			for i := 0; i < leetcode_103_count; i++ {
				temp = append(temp, leetcode_103_collection[i].Val)
			}
			leetcode_103_mutex = false
		} else {
			for i := leetcode_103_count - 1; i >= 0; i-- {
				temp = append(temp, leetcode_103_collection[i].Val)
			}
			leetcode_103_mutex = true
		}

		finalAns := make([]int, len(temp))
		copy(finalAns, temp)
		leetcode_103_ans = append(leetcode_103_ans, finalAns)
		temp = nil

		//已经出了队列的，消失
		leetcode_103_collection = leetcode_103_collection[leetcode_103_count:]
		//下一层的计数器更新
		leetcode_103_current = leetcode_103_future
		//赋予下一层的需要处理的结点数目
		leetcode_103_count = leetcode_103_future

		leetcode_103_future = 0
		//重置处理下标
		leetcode_103_location = 0
	} else {
		//该行还有需要处理的值
		leetcode_103_location++
	}
	//如果下一行没有需要处理的，那就彻底结束了
	if leetcode_103_current == 0 {
		return
	}

	zigzagLevelOrderFunction(leetcode_103_collection[leetcode_103_location])

}

var leetcode_103_ans [][]int

var leetcode_103_collection []*TreeNode

//当前行结点
var leetcode_103_current int

//收集下一行结点
var leetcode_103_future int

//记录好当前行结点的数量
var leetcode_103_count int

//递归传入的下标
var leetcode_103_location int

//反转锁
var leetcode_103_mutex bool
