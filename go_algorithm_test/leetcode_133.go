package main

/*
type Node2 struct {
	Val       int
	Neighbors []*Node2
}

func cloneGraph(node *Node2) *Node2 {
	visited := map[*Node2]*Node2{}
	var cg func(node *Node2) *Node2
	cg = func(node *Node2) *Node2 {
		if node == nil {
			return node
		}

		// 如果该节点已经被访问过了，则直接从哈希表中取出对应的克隆节点返回
		if _, ok := visited[node]; ok {
			return visited[node]
		}

		// 克隆节点，注意到为了深拷贝我们不会克隆它的邻居的列表
		cloneNode := &Node2{node.Val, []*Node2{}}
		// 哈希表存储
		visited[node] = cloneNode

		// 遍历该节点的邻居并更新克隆节点的邻居列表
		for _, node := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(node))
		}
		return cloneNode
	}
	return cg(node)
}

*/

func cloneGraph(node *Node2) *Node2 {
	m := make(map[*Node2]*Node2)

	var res func(node *Node2) *Node2
	//第一次这么写，这表达的意思的内部的函数是什么
	res = func(node *Node2) *Node2 {
		if node == nil {
			return node
		}
		//第一个是val,第二个值是bool型变量，表示是否存在
		if _, ok := m[node]; ok {
			return m[node]
		}
		cloneNode := &Node2{node.Val, []*Node2{}}

		m[node] = cloneNode

		for _, node := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, res(node))
		}

		return cloneNode
	}
	return res(node)
}
