package main

func leetcode_86_partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var ans *ListNode
	var smallPointer *ListNode
	var finalHead *ListNode
	var bigPointer *ListNode
	/*
	 ans 是最终的答案，表示头节点
	 smallPointer 比x小的部分
	 finalHead 使用smallPointer最终指向finalHead使其拼接
	 bigPointer 大于或等于x的部分

	*/
	for ; head != nil; head = head.Next {
		if head.Val < x {
			if ans == nil {
				smallPointer = head
				ans = head
				continue
			}
			smallPointer.Next = head
			smallPointer = smallPointer.Next
		} else {
			if finalHead == nil {
				finalHead = head
				bigPointer = head
				continue
			}
			bigPointer.Next = head
			bigPointer = bigPointer.Next
		}
	}
	//排除掉特殊情况，
	//如果全部都大于等于x
	if ans == nil {
		return finalHead
	}
	smallPointer.Next = finalHead
	//如果全部都小于x
	if bigPointer != nil {
		bigPointer.Next = nil
	}
	return ans
}
