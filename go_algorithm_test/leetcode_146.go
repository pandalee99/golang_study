package main

type LRUCache struct {
	current  int
	capacity int
	head     *LRUNode
	tail     *LRUNode
	search   map[int]*LRUNode
}
type LRUNode struct {
	//应该是一个双向链表
	key      int
	val      int
	previous *LRUNode
	last     *LRUNode
}

func Leetcode_146_Constructor(capacity int) LRUCache {
	return LRUCache{current: 0, capacity: capacity, search: map[int]*LRUNode{}, head: nil, tail: nil}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.search[key]; ok {
		//存在，并更新
		//更新操作================================
		if node == this.head {
			//处于头部则不用改动
		} else if node == this.tail {
			//处于尾部则进行改动，但是注意尾部没有后面的结点
			this.tail = node.previous
			node.previous.last = nil
			node.previous = nil
			node.last = this.head
			this.head.previous = node
			this.head = node
		} else {
			//交叉更改
			node.previous.last = node.last
			node.last.previous = node.previous
			node.previous = nil
			node.last = this.head
			this.head.previous = node
			this.head = node
		}
		//更新操作==================================
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.head == nil {
		//若链表为空，则初始化，并直接返回
		temp := &LRUNode{key, value, nil, nil}
		this.search[key] = temp
		this.head = temp
		this.tail = temp
		this.current++
		return
	}
	//如果已经存在于map
	if node, ok := this.search[key]; ok {
		node.val = value
		//更新操作================================
		if node == this.head {
			//处于头部则不用改动
		} else if node == this.tail {
			//处于尾部则进行改动，但是注意尾部没有后面的结点
			this.tail = node.previous
			node.previous.last = nil
			node.previous = nil
			node.last = this.head
			this.head.previous = node
			this.head = node
		} else {
			//交叉更改
			node.previous.last = node.last
			node.last.previous = node.previous
			node.previous = nil
			node.last = this.head
			this.head.previous = node
			this.head = node
		}
		//更新操作==================================

	} else {
		//不存在于map，添加至头部，并且判断是否溢出
		temp := &LRUNode{key, value, nil, nil}
		this.search[key] = temp
		temp.last = this.head
		this.head.previous = temp
		this.head = temp
		this.current++
		//溢出判断
		if this.current > this.capacity {
			//从map消除指针
			/*
				this.search[this.tail.key] = nil
				这是一个非常错误的做法，实际上并没有将其真正的移除出去，
				正确的做法应该是：delete(this.search,this.tail.key)

			*/
			delete(this.search, this.tail.key)
			//从链表消除指针
			this.tail = this.tail.previous
			if this.tail == nil {
				//说明尾结点为空，表示已经不存在任何结点了，因为刚刚移除的就是唯一的结点
				//于是乎，头结点也要为空
				this.head = nil
			} else {
				this.tail.last = nil
			}
			this.current--
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
