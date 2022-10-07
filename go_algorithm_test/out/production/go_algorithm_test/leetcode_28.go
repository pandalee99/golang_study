package main

func strStr(haystack string, needle string) int {
	//双指针
	n := len(haystack)
	//i是第一指针，needlePointer是第二指针
	needlePointer := 0
	needleLength := len(needle)

	//first表示要返回的初始位置
	first := 0
	for i := 0; i < n; i++ {
		if haystack[i] != needle[needlePointer] {
			if needlePointer != 0 {
				i = first
			}
			needlePointer = 0
		} else {
			if needlePointer == 0 {
				first = i
			}
			needlePointer++
			if needlePointer == needleLength {
				return first
			}
		}

	}

	return -1

}
