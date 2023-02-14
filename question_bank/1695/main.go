package main

// https://leetcode.cn/problems/maximum-erasure-value/
// 1695. 删除子数组的最大得分

// 正整数 删除一个子数组 该子数组含有若干不同元素
// 求子数组和，返回该子数组的最大和
// 不定滑动窗口，窗口性质是，窗口内的所有元素都是不同的
// 选取最大窗口和
// 和之前有一道窗口内没有重复字符的题是一样的思路
// 但是要求和，最佳还是基于前缀和求窗口和
// 为了保证窗口内没有重复的数，我们需要一个 map 或 数组
// 去确认窗口 r 右移时，新进窗口的数是否已经在窗口中，
// 是的话 l 需要右移至窗口内已有的重复数的右侧
func maximumUniqueSubarray(nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, num := range nums {
		sum[i+1] = sum[i] + num
	}
	var l, r, maxSum int
	last := make(map[int]int)
	for r < n {
		// 保证窗口内不存在重复的数
		// 查看新进的数是否已经在窗口，是则右移窗口 l 边界
		if index, ok := last[nums[r]]; ok {
			if index >= l {
				l = index + 1
			}
		}
		// 0 1 3 6 10  现在默认情况下,l,r 都位于下标0
		//   1 2 3 4
		//   0 1 2 3   r=2 l=1
		maxSum = max(maxSum, sum[r+1]-sum[l])
		last[nums[r]] = r
		r++
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
