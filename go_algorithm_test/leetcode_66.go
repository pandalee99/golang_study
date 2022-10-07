package main

func plusOne(digits []int) []int {
	i := -1
	for i = len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	//go切片从头追加数字
	digits = append([]int{1}, digits...)
	return digits
}
