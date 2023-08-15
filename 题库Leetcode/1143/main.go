package main

// https://leetcode.cn/problems/longest-common-subsequence/description/
// 1143. 最长公共子序列

// 最长的公共子序列，显然没办法使用回溯等暴力方式，因为时间复杂度过高，考虑使用 dp
// 最长公共子序列，不需要连续，不存在公共子序列则返回 0
// 定义 dp[i][j] 为 [0~i-1] text1 和以 [0~j-1] text2 的公共子序列的最长长度是多少 (不一定是 text1[i-1]、text2[j-1] 为结尾)
// dp[i][j] = dp[i-1][j-1] + 1, if text1[i-1] == text2[j-1]
// dp[i][j] = max(dp[i-1][j], dp[i][j-1]), if text1[i-1] != text2[j-1]，不相等的话，dp[i][j] 的最长公共子序列可能存在两种情况
// 比如 abcdefg 与 abcpqojygqp, dp[4][4] 显然 text1[4-1] 和 text2[4-1] 不等，此时 dp[4][4] 的取值可能是 dp[4][3] 或 dp[3][4]
func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	var res int
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
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
