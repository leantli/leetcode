package main

// https://leetcode.cn/problems/max-consecutive-ones/
// 485. 最大连续 1 的个数

func findMaxConsecutiveOnes(nums []int) int {
	var maxLen, curLen int
	for _, num := range nums {
		if num == 0 {
			curLen = 0
			continue
		}
		curLen++
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
