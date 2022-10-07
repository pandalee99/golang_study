package main

func minimumTotal(triangle [][]int) int {
	min := 1000
	//行数
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}
	//最后一行的行长
	//m := len(triangle[n-1])
	//初始化
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	//层数为count+1  count表示下标
	count := 1
	for count != n {

		//先处理尾部
		dp[count] = triangle[count][count] + dp[count-1]

		//中间对比部分
		for i := count - 1; i >= 1; i-- {

			if dp[i-1] < dp[i] {
				dp[i] = dp[i-1] + triangle[count][i]
			} else {
				dp[i] = dp[i] + triangle[count][i]
			}

		}
		//后处理头部
		dp[0] = triangle[count][0] + dp[0]

		count++
	}

	//最后找出最小值

	for i := 0; i < n; i++ {
		if dp[i] < min {
			min = dp[i]
		}
	}

	return min
}
