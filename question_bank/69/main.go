package main

// https://leetcode.cn/problems/sqrtx/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 69. x 的平方根

// 这里可以看到，题目要求计算 x 的平方根，但遇到有小数的情况，需要向下取整
// 这里我们可以用 二分，但显然，是要找到某个数的平方数是小于 x 的最大的，比如求 8 的平方数, 2 的平方是 小于 8 的最大值, 3 的平方就大于 8 了
// 算是有一个确切的值，但是可能存在没有该值
func mySqrt(x int) int {
	l, r := 1, x
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	// 出循环后，l 是指向 大于 x 的最小， r 是指向 小于 x 的最大 (mid*mid)
	return r
}
