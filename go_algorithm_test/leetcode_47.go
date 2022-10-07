package main

import (
	"strconv"
)

func permuteUnique(nums []int) [][]int {
	Leetcode_47_ans = nil
	leetcode_47_nums_length = len(nums)
	//要用一个局部的map去判断存不存在
	//把46中本来存放的是值，改为下标，以此达到独一无二的目的
	m := make(map[int]bool)
	//但是即使有了下标，仍然达不到去重的目的，于是乎，我觉得可以再用一个map去解决问题
	leetcode_47_judgmentRepetition = make(map[string]bool)
	var tempAns []int
	for i := 0; i < len(nums); i++ {
		m[i] = true
		targetAns := make([]int, len(tempAns))
		copy(targetAns, tempAns)
		targetAns = append(targetAns, nums[i])
		permuteUniqueRecursiveFunction(m, targetAns, nums, strconv.Itoa(nums[i])+",")
		m[i] = false
	}

	return Leetcode_47_ans
}

func permuteUniqueRecursiveFunction(m map[int]bool, tempAns, nums []int, judge string) {
	if len(tempAns) == leetcode_47_nums_length {
		if leetcode_47_judgmentRepetition[judge] != true {
			leetcode_47_judgmentRepetition[judge] = true
			Leetcode_47_ans = append(Leetcode_47_ans, tempAns)
			return
		}
		return
	}
	for i := 0; i < leetcode_47_nums_length; i++ {
		if m[i] != true {
			m[i] = true
			//tempAns = append(tempAns, nums[i])
			targetAns := make([]int, len(tempAns))
			copy(targetAns, tempAns)
			targetAns = append(targetAns, nums[i])
			permuteUniqueRecursiveFunction(m, targetAns, nums, judge+strconv.Itoa(nums[i])+",")
			m[i] = false
		}
	}
}

var leetcode_47_nums_length int
var Leetcode_47_ans [][]int
var leetcode_47_judgmentRepetition map[string]bool
