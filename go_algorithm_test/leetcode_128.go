package main

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	n := len(nums)
	for i := 0; i < n; i++ {
		m[nums[i]] = true
	}

	ans := 0
	for key, _ := range m {
		//做一个小小的改动
		if m[key-1] == true {
			continue
		}

		i := 1
		max := 1
		for m[key+i] == true {
			max++
			i++
		}

		if max > ans {
			ans = max
		}
	}
	return ans
}
