package main

//
//func connect(root *Node) *Node {
//
//	if root == nil {
//		return nil
//	}
//	//初始化
//	collection = nil
//	current = 1
//	count = 1
//	future = 0
//	location = 0
//	collection = append(collection, root)
//
//	connectFunction(collection[location])
//	return root
//}
//
//func connectFunction(node *Node) {
//	if node == nil {
//		return
//	}
//	//该结点有效
//	current--
//	//并入该结点的子节点
//	if node.Left != nil {
//		future++
//		collection = append(collection, node.Left)
//	}
//	if node.Right != nil {
//		future++
//		collection = append(collection, node.Right)
//	}
//	//主要处理逻辑，该行都处理完了，应该进行下一行的处理
//
//	if current == 0 {
//		node.Next = nil
//		//已经出了队列的，消失
//		collection = collection[count:]
//		//下一层的计数器更新
//		current = future
//		//赋予下一层的需要处理的结点数目
//		count = future
//
//		future = 0
//		//重置处理下标
//		location = 0
//	} else {
//		//该行还有需要处理的值
//		location++
//		node.Next = collection[location]
//	}
//	//如果下一行没有需要处理的，那就彻底结束了
//	if current == 0 {
//		return
//	}
//
//	connectFunction(collection[location])
//}
//
//var collection []*Node
//
////当前行结点
//var current int
//
////收集下一行结点
//var future int
//
////记录好当前行结点的数量
//var count int
//
////递归传入的下标
//var location int
