package main

// https://leetcode.cn/problems/ugly-number/
// 263. 丑数

func isUgly(n int) bool {
	// 丑数是 2^a * 3^b * 5^c
	if n == 0 {
		return false
	}
	temp := []int{2, 3, 5}
	for _, num := range temp {
		for n%num == 0 {
			n /= num
		}
	}
	return n == 1
}
