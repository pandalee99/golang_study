package main

/*
go切片从头追加数字
	var ans []int
	ans = append([]int{1}, ans...)

使用hashmap
	var leetcode_17_map map[int]string
	leetcode_17_map = make(map[int]string)

二维切片初始化
	a := make([][]int, leetcode_17_map) // 二维切片，3行
	for i := range a {
	a[i] = make([]int, n) // 每一行4列

这是一个在二维数组中塞一维数组的方法
ans为二维int数组
				arr := make([]int, 3)
				arr[0] = nums[target]
				arr[1] = nums[first]
				arr[2] = nums[end]
				ans = append(ans, arr)

string[] 切片拼接
	var string_first []string
	string_first = append(string_first, "abc")


得出非矩阵情况下，二维数组的长度
	myArray := [3][4]int{{1,2,3,4},{1,2,3,4},{1,2,3,4}}
    打印一维数组长度
    fmt.Println(len(myArray))
    打印二维数组长度
    fmt.Println(len(myArray[1]))
*/
