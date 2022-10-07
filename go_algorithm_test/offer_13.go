package main

func movingCount(m int, n int, k int) int {
	//先执行对全局变量的初始化操作
	target = k
	leetcode_13_count = 0
	offer_13_length = m
	offer_13_width = n
	offer_13_stepMap = make([][]bool, m) // 二维切片，3行
	for i := range offer_13_stepMap {
		offer_13_stepMap[i] = make([]bool, n) // 每一行4列
	}
	//逻辑
	movingCountSearchFunction(0, 0)
	return leetcode_13_count
}

func movingCountSearchFunction(i int, j int) {
	//墙是不能走的
	if movingCountExist(i, j) == false {
		return
	}
	//走过了就不需要再走了
	if offer_13_stepMap[i][j] == true {
		return
	} else {
		//没走过的路可以踩一脚
		offer_13_stepMap[i][j] = true
		leetcode_13_count++
	}
	//接着就是搜索的问题了
	if j+1 < offer_13_width {
		movingCountSearchFunction(i, j+1)
	}
	if i+1 < offer_13_length {
		movingCountSearchFunction(i+1, j)
	}
	if j-1 >= 0 {
		movingCountSearchFunction(i, j-1)
	}
	if i-1 >= 0 {
		movingCountSearchFunction(i-1, j)
	}
}
func movingCountExist(i int, j int) bool {
	val := 0
	for i >= 10 {
		val += i % 10
		i = i / 10
	}
	val += i
	for j >= 10 {
		val += j % 10
		j = j / 10
	}
	val += j

	if val > target {
		return false
	} else {
		return true
	}

}

var offer_13_stepMap [][]bool

//能走的格子数
var leetcode_13_count int

//墙的判定
var target int

//长和宽
var offer_13_length int
var offer_13_width int
