package main

func combine(n int, k int) [][]int {
	leetcode_77_ans = nil
	for i := 1; i <= n-k+1; i++ {

		tempAns := make([]int, 0)
		tempAns = append(tempAns, i)
		combineRecursiveFunction(tempAns, i+1, n, k-1)
	}

	return leetcode_77_ans
}

func combineRecursiveFunction(tempAns []int, current int, n int, k int) {
	//结束收集
	if k == 0 {
		finalAns := make([]int, len(tempAns))
		copy(finalAns, tempAns)
		leetcode_77_ans = append(leetcode_77_ans, finalAns)
		return
	}
	for i := current; i <= n; i++ {
		targetAns := make([]int, len(tempAns))
		copy(targetAns, tempAns)
		targetAns = append(targetAns, i)
		combineRecursiveFunction(targetAns, i+1, n, k-1)
	}
}

var leetcode_77_ans [][]int
