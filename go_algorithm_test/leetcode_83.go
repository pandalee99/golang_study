package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var first *ListNode
	first = head
	for first.Next != nil {
		for first.Next != nil && first.Val == first.Next.Val {
			first.Next = first.Next.Next
		}
		if first.Next != nil {
			first = first.Next
		} else {
			return head
		}
	}
	return head
}
