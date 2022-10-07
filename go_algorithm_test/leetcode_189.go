package main

/*
最优解应该是翻转三次
*/
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	i := n - k
	j := 0
	var temp []int
	for ; i < n; i++ {
		temp = append(temp, nums[j])
		nums[j] = nums[i]
		j++
	}
	//j=2
	i = 0
	for ; j < n; j++ {
		temp = append(temp, nums[j])
		nums[j] = temp[i]
		i++
	}
}
