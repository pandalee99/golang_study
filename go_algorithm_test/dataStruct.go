package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//type Node struct {
//	Val   int
//	Left  *Node
//	Right *Node
//	Next  *Node
//}

type Node2 struct {
	Val       int
	Neighbors []*Node2
}
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}