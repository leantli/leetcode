package main

// https://leetcode.cn/problems/edit-distance/description/
// 72. 编辑距离

// 双串dp，定义 dp[i][j] 为 word1 0~i-1 区间的单词转换成 word2 0~j-1 区间的单词最少操作多少次
// if word1[i-1] == word2[j-1], dp[i][j] = dp[i-1][j-1]
// if word1[i-1] != word2[j-1], dp[i][j] 的来源有三种，分别是 word1 删除掉 i-1, word2 删除掉 j-1, 替换word1[i-1] 为 word2[j-1]，
// 分别代表着 dp[i-1][j] + 1, dp[i][j-1] + 1, dp[i-1][j-1] + 1
// 因此 dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])+1
// 而初始化则是，dp[0][j] = j, dp[i][0] = i，依照题意和 dp[i][j] 定义进行初始化
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[n][m]
}
func min(args ...int) int {
	minest := args[0]
	for _, arg := range args {
		if arg < minest {
			minest = arg
		}

	}
	return minest
}
