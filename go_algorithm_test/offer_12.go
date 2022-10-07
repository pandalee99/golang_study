package main

func exist(board [][]byte, word string) bool {
	offer_12_flag = false
	var judge [][]bool
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			//开始正式逻辑
			//下标指示器
			k := 0
			//初始化状态图
			judge = make([][]bool, len(board))
			for i := range judge {
				judge[i] = make([]bool, len(board[0]))
			}
			existRecursiveFunction(board, i, j, word, k, judge)
			if offer_12_flag == true {
				return true
			}

		}
	}
	return false
}

func existRecursiveFunction(board [][]byte, i int, j int, word string, k int, judge [][]bool) {
	judge[i][j] = true
	if board[i][j] == word[k] {
		if k+1 == len(word) {
			offer_12_flag = true
			return
		}
		if k+1 > len(word) {
			return
		}
		if i+1 < len(board) {
			if judge[i+1][j] == true {

			} else {
				//judge[i+1][j] = true
				existRecursiveFunction(board, i+1, j, word, k+1, judge)
				judge[i+1][j] = false
			}
		}
		if i-1 >= 0 {
			if judge[i-1][j] == true {

			} else {
				//judge[i-1][j] = true
				existRecursiveFunction(board, i-1, j, word, k+1, judge)
				judge[i-1][j] = false
			}
		}
		if j+1 < len(board[0]) {
			if judge[i][j+1] == true {

			} else {
				//judge[i][j+1] = true
				existRecursiveFunction(board, i, j+1, word, k+1, judge)
				judge[i][j+1] = false
			}
		}
		if j-1 >= 0 {
			if judge[i][j-1] == true {

			} else {
				//judge[i][j-1] = true
				existRecursiveFunction(board, i, j-1, word, k+1, judge)
				judge[i][j-1] = false
			}
		}
	}

}

var offer_12_flag bool
