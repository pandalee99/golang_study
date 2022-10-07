package main

func rotateRight(head *ListNode, k int) *ListNode {
	//排除掉特殊情况
	if k == 0 {
		return head
	}
	if head == nil {
		return nil
	}
	//先保存好指针头的位置
	var first *ListNode
	first = head
	//长度
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	//然后判断k值是否构成了循环,若是，则减去
	//for k > n {
	//	k = k - n
	//}
	//先排除掉一种情况，那就是没有变化的情况
	if k == n {
		return first
	}
	k = n % k
	//这时候的k表示的是，第k个结点指向nil，而第k+1个结点为头结点
	k = n - k
	//然后n也没啥用了，初始化，head指针也一样
	n = 0
	head = first
	//寻找到第k个结点
	for n+1 != k {
		n++
		head = head.Next
	}
	//执行逻辑
	var ans *ListNode
	ans = head.Next
	head.Next = nil
	head = ans
	for head.Next != nil {
		head = head.Next
	}
	head.Next = first

	return ans
}
