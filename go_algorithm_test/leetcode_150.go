package main

import "strconv"

//这题很明显要借助栈去解决问题
//+、-、*、/
func evalRPN(tokens []string) int {
	var stack []int
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" {
			n := len(stack)
			stack[n-2] = stack[n-2] + stack[n-1]
			stack = stack[:n-1]
		} else if tokens[i] == "-" {
			n := len(stack)
			stack[n-2] = stack[n-2] - stack[n-1]
			stack = stack[:n-1]
		} else if tokens[i] == "*" {
			n := len(stack)
			stack[n-2] = stack[n-2] * stack[n-1]
			stack = stack[:n-1]
		} else if tokens[i] == "/" {
			n := len(stack)
			stack[n-2] = stack[n-2] / stack[n-1]
			stack = stack[:n-1]
		} else {
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		}
	}

	return stack[0]
}
