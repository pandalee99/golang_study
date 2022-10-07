package main

/*
不想写数学问题


func rotate(matrix [][]int) {
	//先获得长度
	n := len(matrix)
	//记录旋转次数,分别为当前行的起点和终点，当first>=end则停止
	first, end := 0, n-1
	//j表示行数，二层循环中固定不变,i表示列数
	j := 0
	for first < end {
		for i := first; i < end; i++ {
			//需要临时坐标行 和尾部列
			x, e := j, end
			//试试语法糖
			matrix[x][i], matrix[x][e] = matrix[x][e], matrix[x][i]
		}
		j++
	}

}
*/
