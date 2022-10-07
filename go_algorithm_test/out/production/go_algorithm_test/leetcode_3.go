package main

/*
func lengthOfLongestSubstring(s string) int {
	if s == " " {
		return 1
	}
	if s == "" {
		return 0
	}
	table := map[byte]int{}
	n := len(s)
	result := 0
	count := 0
	for i := 0; i < n; i++ {
		if _, ok := table[s[i]]; ok {
			if count > result {
				result = count
			}
			count = 0
			table = make(map[byte]int)
		}
		table[s[i]] = 1
		count++

	}
	if count > result {
		result = count
	}
	return result
}
*/
//这个方法错了

func lengthOfLongestSubstring(s string) int {

	table := map[byte]int{}
	result, n, start, end := 0, len(s), 0, 0
	for i := 0; i < n; i++ {
		if _, ok := table[s[i]]; ok {
			start = Max(table[s[i]], start)
		}
		result = Max(end-start+1, result)
		table[s[i]] = i + 1
		end++
	}

	return result
}

//第一轮 re=1
//第二轮 re=2
//第三轮 re=2

func Max(nums ...int) int {
	var maxNum = 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}
