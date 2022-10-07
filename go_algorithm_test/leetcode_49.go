package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	count := 0
	judgmentRepetition := make(map[string]int)
	for i := 0; i < len(strs); i++ {
		runeArr := []rune(strs[i])
		sort.Slice(runeArr, func(i, j int) bool {
			return runeArr[i] < runeArr[j] // 正序
		})
		if judgmentRepetition[string(runeArr)] == 0 {
			count++
			judgmentRepetition[string(runeArr)] = count
			ans = append(ans, []string{strs[i]})
		} else {
			ans[judgmentRepetition[string(runeArr)]-1] = append(ans[judgmentRepetition[string(runeArr)]-1], strs[i])
		}
	}

	return ans
}
