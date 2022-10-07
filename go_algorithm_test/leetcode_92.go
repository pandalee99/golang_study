package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	step := 1
	//制作好指针
	var first *ListNode

	//制作好三指针
	var p1 *ListNode
	var p2 *ListNode
	var p3 *ListNode
	p1 = head

	//这个逻辑在于找到p1的前一个位置和p1本身的位置
	if step == left {
		//如果p1就是left，说明没有first结点
		first = nil
	} else {
		for step+1 != left {
			p1 = p1.Next
			step++
		}
		//结束这个循环后,p1要么下一个就是left
		first = p1
		p1 = p1.Next
		step++
	}
	if left == right {
		return head
	}

	p2 = p1.Next
	step++
	if step == right {
		p1.Next = p2.Next
		p2.Next = p1
		if first != nil {
			first.Next = p2
			return head
		} else {
			return p2
		}
	}

	p3 = p2.Next
	step++
	var joinTail *ListNode
	joinTail = p1
	p1.Next = nil
	for step != right {
		p2.Next = p1
		p1 = p2
		p2 = p3
		p3 = p3.Next
		step++
	} //[1,2,3,4,5]
	joinTail.Next = p3.Next
	p2.Next = p1
	p3.Next = p2
	if first == nil {
		return p3
	}
	first.Next = p3

	return head

}
