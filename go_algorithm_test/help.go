package main

/*
二维切片的排序
	slice := [][]int{
		{1, 6},
		{2, 5},
		{3, 6},
	}
	var less func(i, j int) bool
	less = func(i, j int) bool {
		if slice[i][0] == slice[j][0] {
			return slice[i][1] <= slice[j][1]
		}
		return slice[i][0] <= slice[j][0]
	}
	sort.Slice(slice, less)
	fmt.Println(slice)
原理参考
https://zhuanlan.zhihu.com/p/97971489

对字符串排序
	str := "dcdea汉字"
	runeArr := []rune(str)
	sort.Slice(runeArr, func(i, j int) bool {
		return runeArr[i] < runeArr[j] // 正序
	})
	str = string(runeArr)
	fmt.Println(str)

字符串的map
judgmentRepetition := make(map[string]bool)

加快拼接的速度
//需要先导入Strings包
s1 := "字符串"
s2 := "拼接"
var build strings.Builder
build.WriteString(s1)
build.WriteString(s2)
s3 := build.String()

深拷贝
			leetcode_40_ans := make([]int, len(nums))
			copy(leetcode_40_ans ,nums)

为了空间考虑。可以使用byte 去存字符串中的单个值   ''

go切片从头追加数字
	var leetcode_39_ans []int
	leetcode_39_ans = append([]int{1}, leetcode_39_ans...)

使用hashmap
	var leetcode_17_map map[int]string
	leetcode_17_map = make(map[int]string)

二维切片初始化
	a := make([][]int, leetcode_17_map) // 二维切片，3行
	for i := range a {
	a[i] = make([]int, leetcode_46_nums_length) // 每一行4列

这是一个在二维数组中塞一维数组的方法
	ans为二维int数组
				arr := make([]int, 3)
				arr[0] = nums[target]
				arr[1] = nums[first]
				arr[2] = nums[end]
				leetcode_39_ans = append(leetcode_39_ans, arr)

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
