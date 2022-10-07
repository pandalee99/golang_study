package main

func spiralOrder(matrix [][]int) []int {
	var ans []int
	//m行n列
	m := len(matrix)
	n := len(matrix[0])
	//初始化状态图
	judge := make([][]bool, m)
	for i := range judge {
		judge[i] = make([]bool, n)
	}
	//坐标
	x, y := 0, 0
	//第一个值
	judge[x][y] = true
	ans = append(ans, matrix[x][y])
	for true {
		//先往右边走
		for y+1 < n && judge[x][y+1] != true {
			y++
			judge[x][y] = true
			ans = append(ans, matrix[x][y])
		}
		if x+1 == m || judge[x+1][y] == true {
			return ans
		}
		//右边到底了，往下走
		for x+1 < m && judge[x+1][y] != true {
			x++
			judge[x][y] = true
			ans = append(ans, matrix[x][y])
		}
		if y-1 < 0 || judge[x][y-1] == true {
			return ans
		}
		//下边到底了，往左走
		for y-1 >= 0 && judge[x][y-1] != true {
			y--
			judge[x][y] = true
			ans = append(ans, matrix[x][y])
		}
		if x-1 < 0 || judge[x-1][y] == true {
			return ans
		}
		//往上走
		for x-1 >= 0 && judge[x-1][y] != true {
			x--
			judge[x][y] = true
			ans = append(ans, matrix[x][y])
		}
		if y+1 == n || judge[x][y+1] == true {
			return ans
		}
	}

	return ans
}
