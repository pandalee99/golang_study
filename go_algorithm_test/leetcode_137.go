package main

func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, val := range nums {
		m[val] += 1
	}
	for num, val := range m {
		if val == 1 {
			return num
		}
	}
	return 0
}
