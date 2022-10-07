package main

import "strings"

func reverseWords(s string) string {
	var build strings.Builder
	firstMutex := true
	isWordEnd := false
	temp := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && isWordEnd == false {
			continue
		}
		if s[i] == ' ' && isWordEnd {
			build.WriteString(temp)
			temp = ""
			isWordEnd = false
			firstMutex = false
			continue
		}
		if isWordEnd == false {
			isWordEnd = true
			if firstMutex == false {
				build.WriteByte(' ')
			}
		}
		temp = string(s[i]) + temp

	}
	//补充处理
	if isWordEnd {
		build.WriteString(temp)
	}
	return build.String()

}
