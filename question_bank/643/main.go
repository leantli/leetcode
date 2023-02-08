package main

// https://leetcode.cn/problems/maximum-average-subarray-i/
// 643. 子数组最大平均数 I

// 如果是不定长的话，可能会难一些，定长的感觉就没太大难度了
func findMaxAverage(nums []int, k int) float64 {
	// 初始化[0,k) 左闭右开
	l, r := 0, k
	var sum, maxSum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum = sum
	for r < len(nums) {
		sum = sum - nums[l] + nums[r]
		maxSum = max(sum, maxSum)
		l++
		r++
	}
	return float64(maxSum) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
