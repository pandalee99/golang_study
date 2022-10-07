package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	//判断空值
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	//要有一个主串去表示这个逻辑
	mainList := 1
	//给头指针一个坐标
	if list1.Val <= list2.Val {
		head = list1
		mainList = 1

	} else {
		head = list2
		mainList = 2
	}

	var temp *ListNode
	temp = head

	for list1 != nil && list2 != nil {
		if mainList == 1 {
			if list1.Val <= list2.Val {
				temp = list1
				list1 = list1.Next
			} else {
				temp.Next = list2
				mainList = 2

			}
			//1 2 4
			//1 3 4
			//1 1 2 3 4 4

			//1 2 3 6
			//4 5
			//1 2 3 4 5 6
		} else {
			if list2.Val <= list1.Val {
				temp = list2
				list2 = list2.Next
			} else {
				temp.Next = list1
				mainList = 1

			}
		}
	}
	//做收尾处理
	if mainList == 1 {
		temp.Next = list2
	} else {
		temp.Next = list1
	}

	return head
}
