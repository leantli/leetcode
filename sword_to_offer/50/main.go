package main

// https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 50. 第一个只出现一次的字符

// map
// 先不写 map，考虑其他做法
// 位运算？字母可能出现奇数次

// 2. 两次遍历显然不好，考虑同样用 map，但是 value 存的是下标，重复则置为-1, 看看能不能一次遍历
// 想想也显然是不能的，感觉复杂度和法 1 差不多
func firstUniqChar(s string) byte {
	m := make(map[byte]int)
	b := []byte(s)
	for i, v := range b {
		if _, ok := m[v]; ok {
			m[v] = -1
		}
		m[v] = i
	}
	// ... 取不为-1且最小的index，不写了，复杂度差不多
	return ' '
}

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
