package main

import "math"

// https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 42. 连续子数组的最大和

// dp(n) = max{ dp(n-1) + n, n }
// 遍历 dp(n) 的同时选取最大的 dp(n)

func maxSubArray(nums []int) int {
	sum := 0
	max := math.MinInt
	for _, num := range nums {
		sum = Max(sum+num, num)
		max = Max(sum, max)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
