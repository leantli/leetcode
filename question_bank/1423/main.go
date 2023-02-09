package main

// https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards/
// 1423. 可获得的最大点数

// 一时间没有想到这题怎么用滑窗哈哈，明明之前遇到过类似的分区间滑窗思路,刚才居然没想到
// 还可以用定长滑窗，由于我们只取左右两侧的卡牌，显然中间的不取的区间可以作为窗口滑动处理进行计算
// 先算出全部的总和，滑动窗口计算 n-k 长度(即不取的牌)的连续数组的最小和
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	cap := n - k
	var sumW, sumWWMin, sumAll int
	for _, v := range cardPoints[:cap] {
		sumW += v
	}
	sumWWMin = sumW
	sumAll = sumW
	for i := cap; i < n; i++ {
		sumAll += cardPoints[i]
		sumW = sumW - cardPoints[i-cap] + cardPoints[i]
		sumWWMin = min(sumWWMin, sumW)
	}
	return sumAll - sumWWMin
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// // 前缀和+暴力？
// func maxScore(cardPoints []int, k int) int {
// 	n := len(cardPoints)
// 	// 前缀和初始化
// 	sum := make([]int, n+1)
// 	for i, v := range cardPoints {
// 		sum[i+1] = v + sum[i]
// 	}
// 	if n == k {
// 		return sum[n]
// 	}
// 	// 正式暴力
// 	var maxSum int
// 	for i := 0; i <= k; i++ {
// 		maxSum = max(maxSum, sum[i]+(sum[n]-sum[n-k+i]))
// 	}
// 	return maxSum
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
