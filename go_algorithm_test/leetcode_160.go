package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//首先求个各个链表的实际长度
	n, m := 0, 0
	var count *ListNode
	count = headA
	for count != nil {
		count = count.Next
		n++
	}
	count = headB
	for count != nil {
		count = count.Next
		m++
	}

	if n > m {
		n -= m
		for n != 0 {
			headA = headA.Next
			n--
		}
	} else if m > n {
		m -= n
		for m != 0 {
			headB = headB.Next
			m--
		}
	}
	//这样指针对就相等了
	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
