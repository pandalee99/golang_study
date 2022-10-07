package main

func replaceSpace(s string) string {

	var ans string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans += "%20"
			continue
		}
		ans += string(s[i])
	}
	return ans
}
