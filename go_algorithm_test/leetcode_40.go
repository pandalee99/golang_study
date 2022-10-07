package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	leetcode_40_ans = nil
	var currentArray []int
	//不断递归搜索的问题
	for i := 0; i < len(candidates); i++ {
		currentArray = nil
		currentArray = append(currentArray, candidates[i])
		//如果当前就是，那么久不用继续了
		if candidates[i] == target {
			singerAns := []int{candidates[i]}
			leetcode_40_ans = append(leetcode_40_ans, singerAns)
		} else {

			//否则递归性搜索
			combinationSum2RecursiveFunction(candidates, currentArray, i+1, candidates[i], target)
		}
		//防止重复
		for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}
	}
	return leetcode_40_ans
}
func combinationSum2RecursiveFunction(candidates, currentArray []int, begin, sum, target int) {
	for i := begin; i < len(candidates); i++ {
		if candidates[i]+sum == target {
			//不要去改动currentArray
			//等于就加入ans
			//深拷贝
			correctArr := make([]int, len(currentArray))
			copy(correctArr, currentArray)
			//浅拷贝
			//correctArr := currentArray
			correctArr = append(correctArr, candidates[i])
			leetcode_40_ans = append(leetcode_40_ans, correctArr)
		} else if candidates[i]+sum < target {
			//小于则可以直接递归，再次搜索
			smallerArr := currentArray
			smallerArr = append(smallerArr, candidates[i])
			combinationSum2RecursiveFunction(candidates, smallerArr, i+1, candidates[i]+sum, target)
		} else {
			//大于则没有继续搜索的必要了，什么都不做就可以了
		}
		//防止重复
		for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}
	}

}

var leetcode_40_ans [][]int
