package main

func deleteNode(head *ListNode, val int) *ListNode {
	//如果头结点就是，那就简单多了
	if head.Val == val {
		return head.Next
	}
	var ans *ListNode
	ans = head
	//那么得有一个删除用的pointer结点
	var pointer *ListNode
	pointer = head
	for pointer.Next != nil {
		if pointer.Next.Val == val {
			pointer.Next = pointer.Next.Next
			return ans
		}
		pointer = pointer.Next
	}
	return ans
}
