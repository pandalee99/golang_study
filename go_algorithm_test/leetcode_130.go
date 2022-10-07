package main

import "fmt"

func solve(board [][]byte) {
	//初始化
	leetcode_130_m = len(board)
	leetcode_130_n = len(board[0])
	leetcode_130_judge = make(map[point]bool)
	leetcode_130_stay = make(map[point]bool)

	for i := 1; i < leetcode_130_m-1; i++ {
		for j := 1; j < leetcode_130_n-1; j++ {

			p := point{i, j}
			if board[i][j] == 'X' {
				continue
			}
			if leetcode_130_stay[p] {
				continue
			}
			//每次循环前初始化
			leetcode_130_flag = true

			//judge[p] = true
			solveFunction(board, i, j)

			if leetcode_130_flag {
				for p, _ := range leetcode_130_judge {
					board[p.x][p.y] = 'X'
				}
			} else {
				for p, _ := range leetcode_130_judge {
					leetcode_130_stay[p] = true
				}
			}
			leetcode_130_judge = make(map[point]bool)

		}
	}
	for i := 0; i < leetcode_130_m; i++ {
		for j := 0; j < leetcode_130_n; j++ {
			fmt.Print(string(board[i][j]))
		}
		fmt.Println()
	}

}

func solveFunction(board [][]byte, x int, y int) {
	//如果存在碰到边的行为，那么置flag为false ,递归函数结束后直接不处理，并且还要标记
	if y+1 == leetcode_130_n && board[x][y] == 'O' {
		leetcode_130_flag = false
		return
	}
	if x+1 == leetcode_130_m && board[x][y] == 'O' {
		leetcode_130_flag = false
		return
	}

	if y == 0 && board[x][y] == 'O' {
		leetcode_130_flag = false
		return
	}

	if x == 0 && board[x][y] == 'O' {
		leetcode_130_flag = false
		return
	}

	p := point{x, y}
	leetcode_130_judge[p] = true
	if y+1 < leetcode_130_n && board[x][y+1] == 'O' {
		p.y = y + 1
		if leetcode_130_judge[p] == false && leetcode_130_stay[p] == false {
			solveFunction(board, x, y+1)
		}
		p.y = y
	}

	if x+1 < leetcode_130_m && board[x+1][y] == 'O' {
		p.x = x + 1
		if leetcode_130_judge[p] == false && leetcode_130_stay[p] == false {
			solveFunction(board, x+1, y)
		}
		p.x = x
	}

	if y-1 >= 0 && board[x][y-1] == 'O' {
		p.y = y - 1
		if leetcode_130_judge[p] == false && leetcode_130_stay[p] == false {
			solveFunction(board, x, y-1)
		}
		p.y = y
	}

	if x-1 >= 0 && board[x-1][y] == 'O' {
		p.x = x - 1
		if leetcode_130_judge[p] == false && leetcode_130_stay[p] == false {
			solveFunction(board, x-1, y)
		}
		p.x = x
	}

}

type point struct {
	x int
	y int
}

var leetcode_130_judge map[point]bool
var leetcode_130_m int
var leetcode_130_n int
var leetcode_130_flag bool
var leetcode_130_stay map[point]bool
