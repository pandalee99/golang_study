package main

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	collection := map[int]*Node{}
	find := map[*Node]int{}
	i := 0
	//最终的答案头
	var ans *Node
	//跟随指针
	var p *Node
	//重定位指针
	var re *Node
	re = head
	ans = &Node{head.Val, nil, nil}
	p = ans
	collection[i] = ans
	find[head] = i
	i++

	head = head.Next
	for head != nil {
		temp := &Node{head.Val, nil, nil}
		collection[i] = temp
		find[head] = i
		p.Next = temp
		p = p.Next
		head = head.Next
		i++
	}
	//重新定位
	i = 0
	head = re
	p = ans
	for head != nil {
		if head.Random != nil {
			i = find[head.Random]
			p.Random = collection[i]

		}
		head = head.Next
		p = p.Next
	}

	return ans
}

/*
func copyRandomList(head *Node) *Node {
	collection := map[*Node]int{}
	location := 0
	waitMap := map[*Node][]*Node{}
	//最终的答案头
	var ans *Node
	//跟随指针
	var p *Node
	if head == nil {
		return nil
	} else {
		ans = &Node{head.Val, nil, nil}
		p = ans
		collection[ans] = location
		if head.Random != nil {
			val := collection[head.Random]
			if val == 0 {
				ans.Random = ans
			} else {
				temp := &Node{head.Random.Val, nil, nil}
				waitMap[temp] = append(waitMap[temp], ans)
			}
		}

		location++
		head = head.Next
	}

	for head != nil {
		temp := &Node{head.Val, nil, nil}
		p.Next = temp
		collection[temp] = location
		if nodeList, ok := waitMap[temp]; ok {
			for _, node := range nodeList {
				node.Random = temp
			}
		}
		if head.Random != nil {
			val := collection[head.Random]
			if node, ok := collection[temp]; ok {
				temp.Random = node
			} else {
				waitMap[val] = append(waitMap[val], temp)
			}
		}

		//善后
		p = p.Next
		location++
		head = head.Next
	}

	return ans
}
*/
