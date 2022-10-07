package main

import (
	"strconv"
)

func leetcode_9_isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var str string
	str = strconv.Itoa(x)
	head := 0
	tail := len(str) - 1
	for head < tail && head != tail {
		if str[head] != str[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}
