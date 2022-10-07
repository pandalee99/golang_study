package main

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	widthRecord := make(map[int]bool)
	lengthRecord := make(map[int]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				lengthRecord[i] = true
				widthRecord[j] = true
			}
		}
	}

	for x := range lengthRecord {
		for j := 0; j < n; j++ {
			matrix[x][j] = 0
		}
	}
	for x := range widthRecord {
		for i := 0; i < m; i++ {
			matrix[i][x] = 0
		}
	}

}
