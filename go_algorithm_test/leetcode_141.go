package main

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var temp *ListNode
	temp = head.Next

	for head != nil && temp != nil {
		if temp == head {
			return true
		}
		if temp.Next != nil {
			temp = temp.Next
		} else {
			return false
		}
		if temp == head {
			return true
		}
		if temp.Next != nil {
			temp = temp.Next
		} else {
			return false
		}
		if temp == head {
			return true
		}
		head = head.Next
	}

	return false
}
