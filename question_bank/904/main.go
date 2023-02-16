package main

// https://leetcode.cn/problems/fruit-into-baskets/
// 904. 水果成篮

// 只能装两种类型的水果
// 连续子数组 这个数组只能包含两种水果类型
// 求最长的子数组长度
// 不定长滑动窗口
// 窗口性质--窗口中至多有两个不同的整数
func totalFruit(fruits []int) int {
	var l, r, res, curCnt int
	exists := make([]int, len(fruits)+1)
	for r < len(fruits) {
		exists[fruits[r]]++
		if exists[fruits[r]] == 1 {
			curCnt++
		}
		// 当窗口中超过2个不同类型的整数，就右移 l 直至窗口内只剩两种数
		for curCnt > 2 {
			exists[fruits[l]]--
			if exists[fruits[l]] == 0 {
				curCnt--
			}
			l++
		}
		r++
		if res < r-l {
			res = r - l
		}
	}
	return res
}
