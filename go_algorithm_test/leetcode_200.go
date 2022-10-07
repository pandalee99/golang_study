package main

func numIslands(grid [][]byte) int {
	//初始化答案
	ans := 0
	leetcode_200_length = len(grid)
	leetcode_200_width = len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '1' {
				continue
			} else {
				ans++
				//将这块土地包含的1消灭
				numIslandsRecursiveFunction(grid, i, j)
			}
		}
	}

	return ans
}

func numIslandsRecursiveFunction(grid [][]byte, x, y int) {
	grid[x][y] = '0'
	if y+1 < leetcode_200_width && grid[x][y+1] == '1' {
		numIslandsRecursiveFunction(grid, x, y+1)
	}
	if x+1 < leetcode_200_length && grid[x+1][y] == '1' {
		numIslandsRecursiveFunction(grid, x+1, y)
	}
	if y-1 >= 0 && grid[x][y-1] == '1' {
		numIslandsRecursiveFunction(grid, x, y-1)
	}
	if x-1 >= 0 && grid[x-1][y] == '1' {
		numIslandsRecursiveFunction(grid, x-1, y)
	}

}

var leetcode_200_length int
var leetcode_200_width int
