package main

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := 0
	judgeMap := make(map[string]bool)
	judgeMap["1"] = true
	judgeMap["2"] = true
	judgeMap["3"] = true
	judgeMap["4"] = true
	judgeMap["5"] = true
	judgeMap["6"] = true
	judgeMap["7"] = true
	judgeMap["8"] = true
	judgeMap["9"] = true
	judgeMap["10"] = true
	judgeMap["11"] = true
	judgeMap["12"] = true
	judgeMap["13"] = true
	judgeMap["14"] = true
	judgeMap["15"] = true
	judgeMap["16"] = true
	judgeMap["17"] = true
	judgeMap["18"] = true
	judgeMap["19"] = true
	judgeMap["20"] = true
	judgeMap["21"] = true
	judgeMap["22"] = true
	judgeMap["23"] = true
	judgeMap["24"] = true
	judgeMap["25"] = true
	judgeMap["26"] = true

	//其实这道题类似于斐波那契函数

	if len(s) == 1 {
		if judgeMap[string(s[0])] {
			return 1
		} else {
			return 0
		}
	}
	if len(s) == 2 {
		if judgeMap[string(s[0])+string(s[1])] && judgeMap[string(s[1])] {
			return 2
		}
		if judgeMap[string(s[0])+string(s[1])] && judgeMap[string(s[1])] == false {
			return 1
		}
		if judgeMap[string(s[0])] && judgeMap[string(s[1])] {
			return 1
		}
		return 0
	}
	first, second := 0, 1
	if judgeMap[string(s[0])] {
		first = 1
	} else {
		return 0
	}
	if judgeMap[string(s[0])+string(s[1])] && judgeMap[string(s[1])] {
		second = 2
	} else if judgeMap[string(s[0])+string(s[1])] == false && judgeMap[string(s[1])] == false {
		second = 0
	} else {
		second = 1
	}
	for i := 2; i < len(s); i++ {
		if judgeMap[string(s[i])] && judgeMap[string(s[i-1])+string(s[i])] {
			ans = first + second
			first = second
			second = ans
		}
		if judgeMap[string(s[i])] && judgeMap[string(s[i-1])+string(s[i])] == false {
			ans = second
			first = second
			//second = ans
		}
		if judgeMap[string(s[i])] == false && judgeMap[string(s[i-1])+string(s[i])] {
			ans = first
			//first = second
			second = ans
		}
		if judgeMap[string(s[i])] == false && judgeMap[string(s[i-1])+string(s[i])] == false {
			return 0
		}
	}

	return ans
}

/*
	使用递归会超时

func numDecodings(s string) int {

	if len(s) == 0 {
		return 0
	}
	judgeMap = make(map[string]bool)
	judgeMap["1"] = true
	judgeMap["2"] = true
	judgeMap["3"] = true
	judgeMap["4"] = true
	judgeMap["5"] = true
	judgeMap["6"] = true
	judgeMap["7"] = true
	judgeMap["8"] = true
	judgeMap["9"] = true
	judgeMap["10"] = true
	judgeMap["11"] = true
	judgeMap["12"] = true
	judgeMap["13"] = true
	judgeMap["14"] = true
	judgeMap["15"] = true
	judgeMap["16"] = true
	judgeMap["17"] = true
	judgeMap["18"] = true
	judgeMap["19"] = true
	judgeMap["20"] = true
	judgeMap["21"] = true
	judgeMap["22"] = true
	judgeMap["23"] = true
	judgeMap["24"] = true
	judgeMap["25"] = true
	judgeMap["26"] = true

	ans = 0
	n := len(s)
	if n-1 >= 0 && judgeMap[string(s[0])] {
		numDecodingsRecursiveFunction(n-1, s[1:])
	}
	if n-2 >= 0 && judgeMap[string(s[0])+string(s[1])] {
		numDecodingsRecursiveFunction(n-2, s[2:])
	}

	return ans
}

func numDecodingsRecursiveFunction(n int, s string) {
	if n == 0 {
		ans++
	}
	if n-1 >= 0 && judgeMap[string(s[0])] {
		numDecodingsRecursiveFunction(n-1, s[1:])
	}
	if n-2 >= 0 && judgeMap[string(s[0])+string(s[1])] {
		numDecodingsRecursiveFunction(n-2, s[2:])
	}
}

var ans int
var judgeMap map[string]bool
*/
