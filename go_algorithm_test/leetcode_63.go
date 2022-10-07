package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			//能过这条线的 obstacleGrid[i][j]肯定不等于1
			//开阔地带，肯定是自由的
			if i-1 >= 0 && j-1 >= 0 && obstacleGrid[i-1][j] != 1 && obstacleGrid[i][j-1] != 1 {
				dp[j] = dp[j-1] + dp[j]
				continue
			}

			if i-1 >= 0 && j-1 >= 0 && obstacleGrid[i-1][j] == 1 && obstacleGrid[i][j-1] == 1 {
				//如果两边都堵住了，那么你肯定都去不了
				dp[j] = 0
			}
			if i-1 >= 0 && j-1 >= 0 && obstacleGrid[i-1][j] == 1 {
				//因为只有左边这一条路，上面被堵住了，所以dp等于左边
				dp[j] = dp[j-1]
				continue
			}
			if i-1 >= 0 && j-1 >= 0 && obstacleGrid[i][j-1] == 1 {
				//因为只有上面一条路，dp不变，直接continue
				continue
			}
			//初始值
			if i == 0 && j == 0 {
				dp[j] = 1
				continue
			}
			//此时i-1>=0肯定不满足，这是用于第一行的处理
			if j-1 >= 0 {
				dp[j] = dp[j-1]
			}
		}
	}

	return dp[n-1]
}
