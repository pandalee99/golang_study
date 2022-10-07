package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var ans *ListNode
	ans = head.Next
	//双指针法
	var first *ListNode
	var second *ListNode
	var temp *ListNode
	var lastPointer *ListNode
	flag := false
	//
	first = head

	for first != nil {
		//指针位置
		second = first.Next
		if second == nil {
			break
		}
		temp = second.Next
		//首次则不执行
		if flag {
			lastPointer.Next = second

		}
		flag = true
		//记录上一个结点的位置，以保持连续  2,1,4,3,nil
		lastPointer = first
		//交换逻辑
		second.Next = first
		first.Next = temp
		first = temp
	}

	return ans
}
