package main

func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if ans != 0 {
				break
			}
		} else {
			ans++
		}

	}
	return ans
}
