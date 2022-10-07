package main

import "strconv"

func countAndSay(n int) string {
	ans := "1"
	count := 0
	tempAns := ""
	tempSingleNumber := ""

	for i := 1; i < n; i++ {
		for j := 0; j < len(ans); j++ {
			//首次进入
			if tempSingleNumber == "" {
				tempSingleNumber = string(ans[0])
				count++

			} else {
				if tempSingleNumber == string(ans[j]) {
					count++
				} else {
					tempAns += strconv.Itoa(count)
					tempAns += tempSingleNumber
					tempSingleNumber = string(ans[j])
					count = 1
				}
			}
		}
		if tempSingleNumber != "" {
			tempAns += strconv.Itoa(count)
			tempAns += tempSingleNumber
		}
		ans = tempAns
		tempAns = ""
		tempSingleNumber = ""
		count = 0

	}

	return ans
}
