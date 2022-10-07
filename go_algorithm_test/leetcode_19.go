package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//利用三指针法
	var first *ListNode
	var target *ListNode
	var end *ListNode

	first = head
	target = head
	end = head
	//用来判断特殊情况
	flag := 0
	//先让end先走n步  1 5
	for i := 0; i < n; i++ {
		end = end.Next
	}
	//接着便可以找到目标节点了
	for end != nil {
		//则说明至少目标指针走了一次，目标不在头结点处
		if flag == 0 {
			flag = 1
		}
		//如果目标指针还没走，那么头指针就先别走
		if target != head {
			first = first.Next
		}
		target = target.Next
		end = end.Next

	}
	//接着目标指针就到达了目的地
	//说明头指针就是要删除的结点
	if flag == 0 {
		return head.Next
	} else {
		first.Next = target.Next
	}

	return head
}
