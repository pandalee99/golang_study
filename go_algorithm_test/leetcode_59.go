package main

func generateMatrix(n int) [][]int {
	//初始化ans
	var ans [][]int
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	//之后 ans都为0,接着进行循环：
	//坐标
	x, y := 0, 0
	//初始化count
	count := 1
	//赋予首个值
	ans[x][y] = count
	for true {
		//先往右边走
		for y+1 < n && ans[x][y+1] == 0 {
			y++
			count++
			ans[x][y] = count
		}
		if x+1 == n || ans[x+1][y] != 0 {
			return ans
		}
		//右边到底了，往下走
		for x+1 < n && ans[x+1][y] == 0 {
			x++
			count++
			ans[x][y] = count
		}
		if y-1 < 0 || ans[x][y-1] != 0 {
			return ans
		}
		//下边到底了，往左走
		for y-1 >= 0 && ans[x][y-1] == 0 {
			y--
			count++
			ans[x][y] = count
		}
		if x-1 < 0 || ans[x-1][y] != 0 {
			return ans
		}
		//往上走
		for x-1 >= 0 && ans[x-1][y] == 0 {
			x--
			count++
			ans[x][y] = count
		}
		if y+1 == n || ans[x][y+1] != 0 {
			return ans
		}
	}

	return ans
}
