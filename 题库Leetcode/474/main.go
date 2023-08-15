package main

// https://leetcode.cn/problems/ones-and-zeroes/description/
// 474. 一和零

// 暴力显然需要 n^3，其中包含很多重复步骤
// 我认为这里很类似 01 背包，背包的重量上限显然就是 m 个 0 和 n 个 1
// 每个 str 就是每个 物品，0 和 1 的数量都是其重量，价值就是满足的话则可以长度加一
// dp[i][j][k] 表示从 0-i 个 str 中取，放入上限为 j 和 k 的背包中，子集长度最长是多少
// dp[i][j][k] = max(dp[i-1][j-zero[i]][k-one[i]]+1, dp[i-1][j][k]), j>=zero[i] && k >= one[i]
// 再滚动数组优化，减少一维，变成 dp[j][k]，再从后遍历即可 dp[j][k] = max(dp[j-zero[i]][k-one[i]]+1, dp[j][k]), j>=zero[i] && k >= one[i]
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zeroCnt, oneCnt := getZeroAndOneCnt(str)
		for j := m; j >= zeroCnt; j-- {
			for k := n; k >= oneCnt; k-- {
				dp[j][k] = max(dp[j-zeroCnt][k-oneCnt]+1, dp[j][k])
			}
		}
	}
	return dp[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getZeroAndOneCnt(str string) (zero, one int) {
	for _, char := range []byte(str) {
		if char == '1' {
			one++
		} else {
			zero++
		}
	}
	return
}
