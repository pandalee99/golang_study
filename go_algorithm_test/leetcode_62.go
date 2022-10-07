package main

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//逻辑
			if i-1 >= 0 && j-1 >= 0 {
				dp[j] = dp[j-1] + dp[j]
				continue
			}
			//初始值
			if i == 0 && j == 0 {
				dp[j] = 1
				continue
			}
			//此时i-1>=0肯定不满足
			if j-1 >= 0 {
				dp[j] = dp[j-1]
			}
		}
	}

	return dp[n-1]
}

/*//使用dfs 超时，应该使用动态规划！

func uniquePaths(m int, n int) int {
	count = 0
	length = m
	width = n
	stepMap = make([][]bool, m) // 二维切片，m行
	for i := range stepMap {
		stepMap[i] = make([]bool, n) // 每一行n列
	}
	uniquePathsRecursiveFunction(0, 0)
	return count
}

func uniquePathsRecursiveFunction(x int, y int) {
	if x == length-1 && y == width-1 {
		count++
		return
	}
	stepMap[x][y] = true
	if x+1 < length && stepMap[x+1][y] != true {
		uniquePathsRecursiveFunction(x+1, y)
		stepMap[x+1][y] = false
	}
	if y+1 < width && stepMap[x][y+1] != true {
		uniquePathsRecursiveFunction(x, y+1)
		stepMap[x][y+1] = false
	}
}

//长和宽
var length int
var width int
var count int
var stepMap [][]bool


*/
