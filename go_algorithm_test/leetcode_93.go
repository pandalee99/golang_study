package main

import "strconv"

func restoreIpAddresses(s string) []string {

	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	leetcode_93_ans = nil
	number := 0
	if int(s[0]) <= 255 {
		temp := string(s[0]) + "."
		restoreIpAddressesRecursiveFunction(temp, s[1:], 3)
	}
	number, _ = strconv.Atoi(s[:2])
	if number <= 255 && s[0] != '0' {
		temp := string(s[:2]) + "."
		restoreIpAddressesRecursiveFunction(temp, s[2:], 3)
	}
	number, _ = strconv.Atoi(s[:3])
	if number <= 255 && s[0] != '0' {
		temp := string(s[:3]) + "."
		restoreIpAddressesRecursiveFunction(temp, s[3:], 3)
	}

	return leetcode_93_ans
}

func restoreIpAddressesRecursiveFunction(temp string, s string, n int) {

	if s == "" {
		if n == 0 {
			temp = temp[:len(temp)-1]
			leetcode_93_ans = append(leetcode_93_ans, temp)
		}
		return
	}
	if len(s) < n || len(s) > n*3 {
		return
	}
	number := 0
	if int(s[0]) <= 255 {
		//temp += string(s[0]) + "."
		restoreIpAddressesRecursiveFunction(temp+string(s[0])+".", s[1:], n-1)
	}
	if len(s) < 2 {
		return
	}
	number, _ = strconv.Atoi(s[:2])
	if number <= 255 && s[0] != '0' {
		//temp += string(s[:2]) + "."
		restoreIpAddressesRecursiveFunction(temp+string(s[:2])+".", s[2:], n-1)
	}
	if len(s) < 3 {
		return
	}
	number, _ = strconv.Atoi(s[:3])
	if number <= 255 && s[0] != '0' {
		//temp += string(s[:3]) + "."
		restoreIpAddressesRecursiveFunction(temp+string(s[:3])+".", s[3:], n-1)
	}

}

var leetcode_93_ans []string
