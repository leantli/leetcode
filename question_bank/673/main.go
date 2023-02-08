package main

// https://leetcode.cn/problems/number-of-longest-increasing-subsequence/
// 673. 最长递增子序列的个数

// 常规 dp 的基础上，再额外加上计数
// 值得注意的是 计数 防重防漏
// 可能会漏掉的情况，同一个 nums[i] 结尾有多种情况形成最长子序列，然而只计算一次
// 这里计数也要基于 dp 去计数，cnt[i] 表示以 nums[i] 为结尾的最长子序列的个数
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	// dp[i] 表示，以 nums[i] 为结尾的子序列最长的长度
	// cnt[i] 表示，以 nums[i] 为结尾的最长子序列的个数
	// 初始化，默认至少为 1
	dp := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		cnt[i] = 1
	}
	// 初始化，dp[0] 最大长度为 1，此时最长子序列的个数为 1
	maxLen, count := 1, 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// i 前面任意一个 j 元素比 i 小，则可以考虑在 以j为末尾的子序列上补个 i，此时判断长度是否需要变更
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j] // 当前以 nums[i] 为结尾的最长子序列的个数的重置计数
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] == maxLen {
			count += cnt[i]
		} else if dp[i] > maxLen {
			maxLen = dp[i]
			count = cnt[i] // 总体计数的重置计数
		}
	}
	return count
}
