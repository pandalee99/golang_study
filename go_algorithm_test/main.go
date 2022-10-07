package main

import (
	"fmt"
	"sort"
)

func main() {

	test := []int{4, 3, 10, 9, 8}
	//est3 := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa"}
	//test2 := [][]int{{1}}
	//node := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	//node2 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	//test4 := [][]byte{
	//	{'O', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O'},
	//	{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
	//	{'O', 'X', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
	//	{'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O'},
	//	{'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X'},
	//	{'X', 'X', 'O', 'O', 'X', 'O', 'X', 'O', 'X'},
	//	{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
	//	{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
	//	{'O', 'O', 'O', 'O', 'O', 'X', 'X', 'O', 'O'},
	//}
	//test5 := [][]byte{
	//	{'X', 'X', 'X', 'X'},
	//	{'X', 'O', 'O', 'X'},
	//	{'X', 'X', 'O', 'X'},
	//	{'X', 'O', 'X', 'X'},
	//}
	//solve(test5)
	//s := "aaba"
	//intMax = 9223372036854775807
	fmt.Println(minSubsequence(test))

}

func minSubsequence(nums []int) []int {
	var ans []int
	sort.Ints(nums)
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	i := n - 1
	count := nums[i]
	ans = append(ans, nums[i])
	for sum-count >= count {
		i--
		count += nums[i]
		ans = append(ans, nums[i])
	}

	return ans
}
