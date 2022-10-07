package main

func longestCommonPrefix(strs []string) string {

	ans := strs[0]
	if ans == "" {
		return ans
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(ans); j++ {
			if len(ans) > len(strs[i]) && j == len(strs[i]) {
				ans = strs[i]
				break
			}
			if ans[j] != strs[i][j] {
				ans = ans[:j]
				break
			}
		}

		if ans == "" {
			break
		}
	}

	return ans
}
