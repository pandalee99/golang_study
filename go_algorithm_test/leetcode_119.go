package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	count := 0
	var dp []int
	var temp []int
	for count != rowIndex {
		//制作一个数组
		temp = make([]int, count+2)
		//赋予初始值，头部和尾部都为1
		temp[0] = 1
		temp[count] = 1
		for i := 1; i < count; i++ {
			temp[i] = dp[i-1] + dp[i]
		}
		dp = temp
		count++
	}

	return temp
}
