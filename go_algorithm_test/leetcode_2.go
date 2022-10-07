package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//  var	head *ListNode
	//	leetcode_39_ans := ListNode{0, nil}
	//head := ListNode{0, nil} 头部不能这么做

	var head *ListNode
	var tail *ListNode

	//	tempNode := ListNode{0, nil}
	//head.Next = &leetcode_39_ans //头结点
	temp, t1, t2, f := 0, 0, 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			t1 = l1.Val
			l1 = l1.Next
		} else {
			t1 = 0
		}
		if l2 != nil {
			t2 = l2.Val
			l2 = l2.Next
		} else {
			t2 = 0
		}
		temp = t1 + t2 + f
		if temp >= 10 {
			temp = temp - 10
			f = 1
		} else {
			f = 0
		}

		if head == nil {
			head = &ListNode{temp, nil}
			tail = head
		} else {
			tail.Next = &ListNode{temp, nil}
			tail = tail.Next
		}

		////用尾插法建立链表然后返回，链表的头结点
		//tempNode = ListNode{temp, nil}
		//leetcode_39_ans = tempNode
	}
	if f == 1 {
		tail.Next = &ListNode{1, nil}
	}

	return head
}

/*

怎么用尾插法建立链表？

*/
