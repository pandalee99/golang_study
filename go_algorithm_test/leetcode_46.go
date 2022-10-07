package main

func permute(nums []int) [][]int {
	leetcode_46_ans = nil
	leetcode_46_nums_length = len(nums)
	//要用一个局部的map去判断存不存在
	m := make(map[int]bool)
	var tempAns []int
	for i := 0; i < len(nums); i++ {
		m[i] = true
		targetAns := make([]int, len(tempAns))
		copy(targetAns, tempAns)
		targetAns = append(targetAns, nums[i])
		permuteRecursiveFunction(m, targetAns, nums)
		m[i] = false
	}

	return leetcode_46_ans
}

func permuteRecursiveFunction(m map[int]bool, tempAns, nums []int) {
	if len(tempAns) == leetcode_46_nums_length {
		leetcode_46_ans = append(leetcode_46_ans, tempAns)
		return
	}
	for i := 0; i < leetcode_46_nums_length; i++ {
		if m[i] != true {
			m[i] = true
			//tempAns = append(tempAns, nums[i])
			targetAns := make([]int, len(tempAns))
			copy(targetAns, tempAns)
			targetAns = append(targetAns, nums[i])
			permuteRecursiveFunction(m, targetAns, nums)
			m[i] = false
		}
	}
}

var leetcode_46_nums_length int
var leetcode_46_ans [][]int
