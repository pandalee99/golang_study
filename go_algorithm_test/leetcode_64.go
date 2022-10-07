package main

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//逻辑
			if i-1 >= 0 && j-1 >= 0 {
				if dp[j] < dp[j-1] {
					dp[j] = grid[i][j] + dp[j]
				} else {
					dp[j] = grid[i][j] + dp[j-1]
				}
				continue
			}
			//初始值
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
				continue
			}
			//此时i-1>=0肯定不满足
			if j-1 >= 0 {
				dp[j] = grid[i][j] + dp[j-1]
			} else {
				dp[j] = grid[i][j] + dp[j]
			}
		}
	}

	return dp[n-1]
}
