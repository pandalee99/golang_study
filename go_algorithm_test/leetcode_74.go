package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)

	for i := 0; i < m; i++ {
		if matrix[i][0] == target {
			return true
		}

		if matrix[i][0] < target && i+1 != m {
			continue
		}
		//那么就只剩下matrix[i][0]>target
		//二分查找
		//排除掉特殊情况
		if matrix[i][0] > target && i-1 == -1 {
			return false
		}
		//排除掉在边界位置的情况
		if matrix[i][0] < target && i+1 == m {
			i++
		}

		first := 0
		end := len(matrix[0]) - 1
		for first <= end {
			mid := (first + end) / 2
			if matrix[i-1][mid] == target {
				return true
			} else if matrix[i-1][mid] > target {
				end = mid - 1
			} else {
				first = mid + 1
			}
		}
		return false

	}

	return false
}
