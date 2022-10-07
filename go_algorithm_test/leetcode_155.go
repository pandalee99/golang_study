package main

type MinStack struct {
	list    []int
	minList []int
}

func Leetcode_155_Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(val int) {
	this.list = append(this.list, val)
	if len(this.minList) == 0 {
		this.minList = append(this.minList, val)
	} else {
		if val < this.minList[len(this.minList)-1] {
			this.minList = append(this.minList, val)
		} else {
			this.minList = append(this.minList, this.minList[len(this.minList)-1])
		}
	}
}

func (this *MinStack) Pop() {
	this.list = this.list[:len(this.list)-1]
	this.minList = this.minList[:len(this.minList)-1]
}

func (this *MinStack) Top() int {
	return this.list[len(this.list)-1]
}

func (this *MinStack) GetMin() int {
	return this.minList[len(this.minList)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/*


type MinStack struct {
	list []*minStackListNode
	min  *minStackListNode
}
type minStackListNode struct {
	val  int
	pre  *minStackListNode
	next *minStackListNode
}

func Constructor() MinStack {
	return MinStack{[]*minStackListNode{}, nil}
}

func (this *MinStack) Push(val int) {
	node := &minStackListNode{val: val}
	this.list = append(this.list, node)
	if this.min == nil {
		this.min = node
	} else {
		//如果不等于空，那么看看如何进行连接

		//如果就是最小了，直接变为新的头部
		if node.val <= this.min.val {
			node.next = this.min
			this.min.pre = node
			this.min = node
		} else {
			//不是最小，那么寻找合适的位置放入
			p := this.min
			for p.next != nil && p.val <= node.val {
				p = p.next
			}
			//进行完了这个循环后，有两种可能：
			if p.val < node.val {
				//node是最大的
				p.next = node
				node.pre = p
			} else {
				//或者是p.val>node.val 进行新的链接
				node.next = p
				node.pre = p.pre

				p.pre.next = node
				p.pre = node
			}
		}
	}

}

func (this *MinStack) Pop() {
	node := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	//如果没了，直接清空
	if len(this.list) == 0 {
		this.min = nil
	} else {
		//至少还有一个，则看看链表的处理

		//如果刚刚去掉的结点就是最小结点
		if node == this.min {
			this.min = this.min.next
			this.min.pre = nil
		} else {
			//不是最小的结点

			if node.next == nil {
				//但是这个在最后面,去掉的结点是最大的结点
				node.pre.next = nil
				node.pre = nil
			} else {
				//在中间
				node.pre.next = node.next
				node.next.pre = node.pre
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.list[len(this.list)-1].val
}

func (this *MinStack) GetMin() int {
	return this.min.val
}


*/
