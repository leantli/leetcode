package main

// https://leetcode.cn/problems/uncrossed-lines/
// 1035. 不相交的线

// 相当于是在 nums1 和 nums2 中选取相等的数，其中保证后面的选取的数，i 和 j 都大于之前选取过的
// 然后求出相等数最多的数量，其实就是类似于求 最长的公共子序列 T1143
// dp[i][j] 定义为 nums1 中 0~i-1 下标，和nums2 中 0~j-1 下标，不相交的线的数量最多是多少
// dp[i][j] = dp[i-1][j-1] + 1, if nums1[i-1] == nums2[j-1]
// dp[i][j] = max(dp[i-1][j], dp[i][j-1]), if nums1[i-1] != nums2[j-1]
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	var res int
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
			res = max(res, dp[i][j])
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
