package main

import "sort"

// https://leetcode.cn/problems/pile-box-lcci/
// 面试题 08.13. 堆箱子

// 这是道 LIS 三维属性的题目，我们无法像二维时
// 一维升序，二维降序，进而简单使用 LIS dp
// 但是我们可以通过多加判断保证状态转移无误，当然，这可能也是最普通最暴力的dp
func pileBox(box [][]int) int {
	// 一维升序，一维等时二维升，一二等时三维升
	sort.Slice(box, func(i, j int) bool {
		return box[i][0] < box[j][0] || (box[i][0] == box[j][0] && box[i][1] < box[j][1]) || (box[i][0] == box[j][0] && box[i][1] == box[j][1] && box[i][2] < box[j][2])
	})
	// dp[i] 表示 以 box[i] 作为最底部箱子时，箱堆的最大高度
	dp := make([]int, len(box))
	// 初始化，所有箱子都能以自己为单独的箱堆，至少高度都为1
	for i := range box {
		dp[i] = box[i][2]
	}
	res := dp[0]
	for i := 1; i < len(box); i++ {
		for j := 0; j < i; j++ {
			if box[j][0] < box[i][0] && box[j][1] < box[i][1] && box[j][2] < box[i][2] {
				dp[i] = max(dp[i], dp[j]+box[i][2])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
