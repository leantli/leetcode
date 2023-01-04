package main

// https://leetcode.cn/problems/valid-perfect-square/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 367. 有效的完全平方数

// 二分解？
// 有准确值否？算是有吧，采用准确值模板
func isPerfectSquare(num int) bool {
	l, r := 1, num
	for l <= r {
		mid := l + (r-l)/2
		temp := mid * mid
		if temp == num {
			return true
		} else if temp > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
