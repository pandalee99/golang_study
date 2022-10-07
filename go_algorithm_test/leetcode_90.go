package main

import (
	"sort"
	"strconv"
)

func subsetsWithDup(nums []int) [][]int {
	//初始化
	leetcode_90_ans = nil
	leetcode_90_judgeMap = make(map[string]bool)
	sort.Ints(nums)
	leetcode_90_ans = append(leetcode_90_ans, nil)
	leetcode_90_globeNums = nums

	for j := 1; j <= len(nums); j++ {
		judgeAns := ""

		for i := 0; i < len(nums)-j+1; i++ {
			if i > 0 && nums[i-1] == nums[i] {
				continue
			}
			tempAns := make([]int, 0)
			tempAns = append(tempAns, nums[i])
			subsetsWithDupRecursiveFunction(tempAns, judgeAns, j-1, nums[i], i+1)
		}

	}
	return leetcode_90_ans
}

func subsetsWithDupRecursiveFunction(tempAns []int, judgeAns string, n int, target int, begin int) {
	judgeAns += strconv.Itoa(target) + ","
	if n == 0 {
		if leetcode_90_judgeMap[judgeAns] == false {
			leetcode_90_judgeMap[judgeAns] = true
			leetcode_90_ans = append(leetcode_90_ans, tempAns)
		}
		return
	}

	for i := begin; i < len(leetcode_90_globeNums); i++ {
		copyAns := make([]int, len(tempAns))
		copy(copyAns, tempAns)
		copyAns = append(copyAns, leetcode_90_globeNums[i])
		subsetsWithDupRecursiveFunction(copyAns, judgeAns, n-1, leetcode_90_globeNums[i], i+1)
	}

}

var leetcode_90_ans [][]int
var leetcode_90_judgeMap map[string]bool
var leetcode_90_globeNums []int
