package main

// https://leetcode.cn/problems/squares-of-a-sorted-array/
// 977. 有序数组的平方

// 数组非递减
// 返回每个数字的平方组成的新数组
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l, r := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if compare(nums[l], nums[r]) {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
	}
	return res
}

// a>b 返回 true
func compare(a, b int) bool {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a > b
}
