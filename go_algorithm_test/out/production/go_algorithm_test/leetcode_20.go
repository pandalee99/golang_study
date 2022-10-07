package main

func isValid(s string) bool {
	//模拟一个栈
	//'('，')'，'{'，'}'，'['，']'
	var stack string
	top := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack += "("
			top++
		}
		if s[i] == '[' {
			stack += "["
			top++
		}
		if s[i] == '{' {
			stack += "{"
			top++
		}
		//开始进行一一匹配
		if s[i] == ')' {
			if top == 0 {
				return false
			}
			if stack[top-1] == '(' {
				top--
				stack = stack[:top]
			} else {
				return false
			}
		}
		if s[i] == ']' {
			if top == 0 {
				return false
			}
			if stack[top-1] == '[' {
				top--
				stack = stack[:top]
			} else {
				return false
			}
		}
		if s[i] == '}' {
			if top == 0 {
				return false
			}
			if stack[top-1] == '{' {
				top--
				stack = stack[:top]
			} else {
				return false
			}
		}

	}

	if top == 0 {
		return true
	}

	return false
}
