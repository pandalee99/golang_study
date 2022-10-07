package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	var temp *ListNode
	var preNode *ListNode
	temp = head.Next
	preNode = head

	for temp != nil {
		if temp.Val <= head.Val {
			change := temp

			//移除结点
			temp = temp.Next
			preNode.Next = temp
			//移除结点

			change.Next = head
			head = change
		} else if temp.Val >= preNode.Val {
			preNode = preNode.Next
			temp = temp.Next
		} else {
			//temp.Val > head.Val
			change := temp
			p := head

			//移除结点
			temp = temp.Next
			preNode.Next = temp
			//移除结点

			for p.Next.Val < change.Val {
				p = p.Next
			}
			//出了这个循环，只会是指针p.next.val>=change
			//那么p本身是小于change的
			change.Next = p.Next
			p.Next = change
		}

	}

	return head
}
