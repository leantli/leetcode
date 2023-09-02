package main

import "strconv"

// https://leetcode.cn/problems/add-strings/description/
// 415. 字符串相加

// 从后往前遍历，需要一个 upper 存储进位的值
// 使用 i, j 作为两个字符串的指针，当某个指针越界，则对应的数字直接给定为 0，不需要读取进行转换
func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	var res string
	var upper int
	for i >= 0 || j >= 0 || upper != 0 {
		var n1, n2 int
		if i >= 0 {
			n1, _ = strconv.Atoi(string(num1[i]))
		}
		if j >= 0 {
			n2, _ = strconv.Atoi(string(num2[j]))
		}
		sum := n1 + n2 + upper
		res = strconv.Itoa(sum%10) + res
		upper = sum / 10
		i--
		j--
	}
	return string(res)
}
