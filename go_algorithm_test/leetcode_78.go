package main

func subsets(nums []int) [][]int {
	leetcode_78_ans = nil
	leetcode_78_nums = nums
	leetcode_78_ans = append(leetcode_78_ans, nil)
	for i := 1; i <= len(nums); i++ {
		//这里i表示框的大小
		for j := 1; j <= len(nums)-i+1; j++ {
			tempAns := make([]int, 0)
			tempAns = append(tempAns, leetcode_78_nums[j-1])
			subsetsRecursiveFunction(tempAns, j+1, len(nums), i-1)
		}

	}

	return leetcode_78_ans
}

func subsetsRecursiveFunction(tempAns []int, current int, n int, k int) {
	//结束收集
	if k == 0 {
		finalAns := make([]int, len(tempAns))
		copy(finalAns, tempAns)
		leetcode_78_ans = append(leetcode_78_ans, finalAns)
		return
	}
	for i := current; i <= n; i++ {
		targetAns := make([]int, len(tempAns))
		copy(targetAns, tempAns)
		targetAns = append(targetAns, leetcode_78_nums[i-1])
		subsetsRecursiveFunction(targetAns, i+1, n, k-1)
	}
}

var leetcode_78_ans [][]int
var leetcode_78_nums []int
