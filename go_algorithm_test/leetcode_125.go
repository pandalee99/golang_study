package main

func leetcode_125_isPalindrome(s string) bool {

	n := len(s)
	i, j := 0, n-1
	for i < j {
		for isPalindromeHelper(s[i]) {
			i++
			if i == j {
				return true
			}
		}
		for isPalindromeHelper(s[j]) {
			j--
			if i == j {
				return true
			}
		}
		val1 := s[i]
		if s[i] <= 122 && s[i] >= 97 {
			val1 -= 32
		}
		val2 := s[j]
		if s[j] <= 122 && s[j] >= 97 {
			val2 -= 32
		}
		if val1 != val2 {
			return false
		}
		i++
		j--
	}

	return true
}

func isPalindromeHelper(u uint8) bool {
	if (u <= 122 && u >= 97) || (u <= 90 && u >= 65) || (u <= 57 && u >= 48) {
		return false
	}
	return true
}
