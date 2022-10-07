package main

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	//-10,-3,0,5,9
	target := n / 2
	val := nums[target]
	ans := &TreeNode{val, nil, nil}
	if target > 0 {
		ans.Left = sortedArrayToBST(nums[:target])
	}
	if target < n-1 {
		ans.Right = sortedArrayToBST(nums[target+1:])
	}

	return ans
}
