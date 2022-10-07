package main

func generateParenthesis(n int) []string {

	//全局变量初始化
	flag = true
	leetcode_22_ans = nil
	//一眼递归
	//利用虚拟栈和实际栈去解决问题，空间复杂度会很大
	generateParenthesisRecursiveFunctionLeft("", "", 0, n, n)
	generateParenthesisRecursiveFunctionRight("", "", 0, n, n)

	return leetcode_22_ans
}

func generateParenthesisRecursiveFunctionLeft(tempAns, stack string, top, left, right int) {
	if left == 0 && right == 0 {
		if flag == true {
			flag = false
			leetcode_22_ans = append(leetcode_22_ans, tempAns)
			return
		} else {
			flag = true
		}

	}
	if right < left {
		return
	}
	stack += "("
	tempAns += "("
	top++
	left--
	if right >= 0 && left >= 0 {
		generateParenthesisRecursiveFunctionLeft(tempAns, stack, top, left, right)
		generateParenthesisRecursiveFunctionRight(tempAns, stack, top, left, right)
	}

}
func generateParenthesisRecursiveFunctionRight(tempAns, stack string, top, left, right int) {
	if left == 0 && right == 0 {
		if flag == true {
			flag = false
			leetcode_22_ans = append(leetcode_22_ans, tempAns)
			return
		} else {
			flag = true
		}

	}
	if top == 0 {
		return
	}
	if stack[top-1] == '(' {
		top--
		stack = stack[:top]
		tempAns += ")"
		right--
	}
	if right >= 0 && left >= 0 {
		generateParenthesisRecursiveFunctionLeft(tempAns, stack, top, left, right)
		generateParenthesisRecursiveFunctionRight(tempAns, stack, top, left, right)
	}
}

var leetcode_22_ans []string
var flag = true

/*
func generateParenthesisRecursiveFunction(tempAns, stack string, top, left, right int) {
	if left > 0 {
		if right < left {
			return
		}

		stack += "("
		tempAns += "("
		top++
		left--
		if left == 0 {

		} else {
			generateParenthesisRecursiveFunction(tempAns, stack, top, left, right)
		}
	}
	if right > 0 {
		if top == 0 {
			return
		}
		if stack[top-1] == '(' {
			top--
			stack = stack[:top]
			tempAns += ")"
			generateParenthesisRecursiveFunction(tempAns, stack, top, left, right-1)
		} else {
			return
		}
	}
	if left == 0 && right == 0 {
		leetcode_22_ans = append(leetcode_22_ans, tempAns)
		return
	}
}
*/
