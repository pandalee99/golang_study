package main

func findNumberIn2DArray(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}
	n := len(matrix)
	m := len(matrix[0])
	j := 0
	looked := 0
	for i := 0; i < n; i++ {
		for j < m && j > -1 {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] < target {
				looked = matrix[i][j]
				j++
				continue
			}
			if matrix[i][j] > target {
				//注意边界
				if j-1 > -1 {
					j--
				} else {
					break
				}

				if matrix[i][j] == looked {
					break
				}
				//break
			}
		}
		//边界条件
		if j == m {
			j--
		}
	}

	return false
}
