package main

/*
解法错误、
*/
func twoSum(numbers []int, target int) []int {
	//n := len(numbers)
	i := len(numbers) - 1
	//for i = n - 1; i >= 0; i-- {
	//	//如果这个数已经大于了，那么肯定是不可能的
	//	if numbers[i] >= target {
	//		continue
	//	} else {
	//		break
	//	}
	//}
	//1,4,5,7,8,9,10,11,22,23,24,25   6
	//声明
	a, b := i-1, i
	if numbers[a]+numbers[b] == target {
		return []int{a + 1, b + 1}
	}
	//如果没有返回，那么肯定不是这两个
	for i = a - 1; i >= 0; i-- {
		if numbers[i]+numbers[a] == target {
			return []int{i + 1, a + 1}
		} else if numbers[i]+numbers[a] > target {
			b = a
			a = i
			continue
		} else {
			//如果发生了i+a< 但i+b>
			//那么，b肯定是最终答案的一个答案
			//如此一来，只需要搜索到正确的i就可了
			i++
			for numbers[i]+numbers[b] != target {
				i--
			}
			return []int{i + 1, b + 1}
		}
	}
	return nil
}
