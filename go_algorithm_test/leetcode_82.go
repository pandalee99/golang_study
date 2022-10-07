package main

func Leetcode_82_deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//保证头部指针不发生重复现象
	last := head.Val
	for head.Next != nil && last == head.Next.Val {
		head = head.Next
		if head.Next == nil {
			return nil
		}
		for last == head.Val {
			if head.Next == nil {
				return nil
			}
			head = head.Next
		}
		last = head.Val
	} //
	//凡是final指针指示的值都绝对是不重复的
	var final *ListNode
	final = head
	var first *ListNode
	var second *ListNode
	//指针初始化
	if final.Next != nil {
		second = final.Next
	} else {
		return head
	}
	if second.Next != nil {
		first = second.Next
	} else {
		return head
	}
	//不能忽略奇数的情况
	last = second.Val
	//开始是否重复逻辑的判断
	for first.Next != nil || second.Val == first.Val {
		if second.Val != first.Val { //[1,2,2,3,3]
			//全部前进一步
			final = final.Next
			second = second.Next
			first = first.Next
			last = second.Val
		} else {
			//说明second指针和first指针相等
			final.Next = first.Next
			//重新开始指针初始化
			if final.Next != nil {
				second = final.Next
			} else {
				return head
			}
			//当为3 的时候，需要跳跃判断
			for second.Val == last {
				if second.Next == nil {
					final.Next = nil
					return head
				} else {
					second = second.Next
					final.Next = second
				}
			}
			if second.Next != nil {
				first = second.Next
			} else {
				return head
			}

			//不能忽略奇数的情况
			last = second.Val
		}
	}

	return head
}
