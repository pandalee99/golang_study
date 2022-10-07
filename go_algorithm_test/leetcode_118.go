package main

func generate(numRows int) [][]int {
	var ans [][]int
	ans = append(ans, []int{1})
	if numRows == 1 {
		return ans
	}
	count := 1
	var dp []int
	for count != numRows {
		//制作一个数组
		temp := make([]int, count+1)
		//赋予初始值，头部和尾部都为1
		temp[0] = 1
		temp[count] = 1
		for i := 1; i < count; i++ {
			temp[i] = dp[i-1] + dp[i]
		}
		dp = temp
		ans = append(ans, temp)
		count++
	}

	return ans
}
