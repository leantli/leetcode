package main

// https://leetcode.cn/problems/video-stitching/
// 1024. 视频拼接

// 目标：最小数目的剪辑片段，达到覆盖time [0, time] 的长度
// dp[i] 定义为 覆盖 i 长度 [0,i) 的片段最少需要多少剪辑片段?
// 那么怎么定义其状态转移方程？
// dp[0] = 0，dp[i] = min(dp[i], dp[j]+1), j 如何指代？
// 当 x 剪辑视频为 [l,r) 时，当 l < i <= r 时，dp[i] = min(dp[i], dp[l]+1)
// 这里我们尽量从 l 下标取值，越往左，需要的剪辑数量要么相等要么越少
func videoStitching(clips [][]int, time int) int {
	maxCnt := 111 // 大于 100 即可，题目给出的剪辑片段最多为 100
	dp := make([]int, time+1)
	for i := 1; i <= time; i++ {
		dp[i] = maxCnt
	}
	for i := 1; i <= time; i++ {
		for _, c := range clips {
			l, r := c[0], c[1]
			if l < i && i <= r && dp[i] > dp[l]+1 {
				dp[i] = dp[l] + 1
			}
		}
	}
	if dp[time] == maxCnt {
		return -1
	}
	return dp[time]
}
