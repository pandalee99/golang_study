package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	//排除特殊情况
	if head.Next == nil || head.Next.Next == nil {
		return
	}
	var temp *ListNode
	var p *ListNode
	temp = head
	p = head
	for temp != nil {
		temp = temp.Next
		if temp != nil {
			temp = temp.Next
		} else {
			break
		}
		p = p.Next
	}

	//这段执行完后p就是头部了，然后开始反转链表
	temp = p.Next
	var first *ListNode
	if temp.Next != nil {
		first = temp.Next
		p.Next = nil
		for first != nil {
			temp.Next = p
			p = temp
			temp = first
			first = first.Next
		}
		temp.Next = p
		first = temp

	} else {
		first = temp
		first.Next = p
		p.Next = nil
	}
	//之后就变成了 head和first 两个头部指针了
	temp = nil
	p = nil
	for true {
		temp = head.Next
		head.Next = first
		head = temp
		if temp.Next == nil {
			break
		}
		p = first.Next
		first.Next = head
		first = p
	}
}
