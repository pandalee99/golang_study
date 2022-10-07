package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	temp := head
	for k != 0 {
		temp = temp.Next
		k--
	}
	for temp != nil {
		head = head.Next
		temp = temp.Next
	}

	return head
}
