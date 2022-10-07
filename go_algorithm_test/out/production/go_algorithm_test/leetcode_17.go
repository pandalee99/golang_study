package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	//使用全局变量的时候需要初始化一下
	leetcode_17_ans = nil
	leetcode_17_map = make(map[int]string)
	leetcode_17_map[2] = "abc"
	leetcode_17_map[3] = "def"
	leetcode_17_map[4] = "ghi"
	leetcode_17_map[5] = "jkl"
	leetcode_17_map[6] = "mno"
	leetcode_17_map[7] = "pqrs"
	leetcode_17_map[8] = "tuv"
	leetcode_17_map[9] = "wxyz"

	letterCombinationsRecursiveFunction(digits, len(digits), "")
	return leetcode_17_ans
}
func letterCombinationsRecursiveFunction(digits string, count int, perAns string) {
	if count == 0 {
		//拼接
		leetcode_17_ans = append(leetcode_17_ans, perAns)

	} else {
		key := int(digits[0]) - 48     //uint 转 int
		n := len(leetcode_17_map[key]) //该字符串的长度
		for i := 0; i < n; i++ {

			letterCombinationsRecursiveFunction(digits[1:], count-1, perAns+string(leetcode_17_map[key][i]))

		}
	}

}

var leetcode_17_ans []string
var leetcode_17_map map[int]string
