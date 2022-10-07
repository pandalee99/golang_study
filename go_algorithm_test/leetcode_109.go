package main

func sortedListToBST(head *ListNode) *TreeNode {

	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	if nums == nil {
		return nil
	}
	return sortedListToBSTHelperFunction(nums)
}

func sortedListToBSTHelperFunction(nums []int) *TreeNode {
	n := len(nums)
	//-10,-3,0,5,9
	target := n / 2
	val := nums[target]
	ans := &TreeNode{val, nil, nil}
	if target > 0 {
		ans.Left = sortedListToBSTHelperFunction(nums[:target])
	}
	if target < n-1 {
		ans.Right = sortedListToBSTHelperFunction(nums[target+1:])
	}

	return ans
}
