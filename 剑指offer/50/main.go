package main

// https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 50. 第一个只出现一次的字符

// map
// 先不写 map，考虑其他做法
// 位运算？字母可能出现奇数次

// 1. 先写 map 吧，两次遍历
// func firstUniqChar(s string) byte {
// 	m := make(map[byte]int)
// 	b := []byte(s)
// 	for _, v := range b {
// 		m[v]++
// 	}
// 	for _, v := range b {
// 		if m[v] == 1 {
// 			return v
// 		}
// 	}
// 	return ' '
// }
