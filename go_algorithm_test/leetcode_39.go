package main

func combinationSum(candidates []int, target int) [][]int {
	leetcode_39_ans = nil
	var currentArray []int
	//不断递归搜索的问题
	for i := 0; i < len(candidates); i++ {
		currentArray = nil
		currentArray = append(currentArray, candidates[i])
		//如果当前就是，那么久不用继续了
		if candidates[i] == target {
			singerAns := []int{candidates[i]}
			leetcode_39_ans = append(leetcode_39_ans, singerAns)
		} else {

			//否则递归性搜索
			combinationSumRecursiveFunction(candidates, currentArray, i, candidates[i], target)
		}
	}
	return leetcode_39_ans
}
func combinationSumRecursiveFunction(candidates, currentArray []int, begin, sum, target int) {
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
			leetcode_39_ans = append(leetcode_39_ans, correctArr)
		} else if candidates[i]+sum < target {
			//小于则可以直接递归，再次搜索
			smallerArr := currentArray
			smallerArr = append(smallerArr, candidates[i])
			combinationSumRecursiveFunction(candidates, smallerArr, i, candidates[i]+sum, target)
		} else {
			//大于则没有继续搜索的必要了，什么都不做就可以了
		}
	}

}

var leetcode_39_ans [][]int
