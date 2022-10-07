package main

//type CQueue struct {
//	inStack, outStack []int
//}
//
//func Constructor() CQueue {
//	return CQueue{}
//}
//
//func (this *CQueue) AppendTail(value int) {
//	this.inStack = append(this.inStack, value)
//
//}
//
//func (this *CQueue) DeleteHead() int {
//	if len(this.outStack) == 0 {
//		if len(this.inStack) == 0 {
//			return -1
//		}
//		this.in2out()
//	}
//	value := this.outStack[len(this.outStack)-1]
//	this.outStack = this.outStack[:len(this.outStack)-1]
//	return value
//}
//func (this *CQueue) in2out() {
//	for len(this.inStack) != 0 {
//		value := this.inStack[len(this.inStack)-1]
//		this.outStack = append(this.outStack, value)
//		this.inStack = this.inStack[:len(this.inStack)-1]
//	}
//}

/*
解法一,使用slice

type CQueue struct {
	QueueInt []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.QueueInt = append(this.QueueInt, value)

}

func (this *CQueue) DeleteHead() int {
	if len(this.QueueInt) == 0 {
		return -1
	} else {
		leetcode_39_ans := this.QueueInt[0]
		this.QueueInt = this.QueueInt[1:]
		return leetcode_39_ans
	}

}
*/
