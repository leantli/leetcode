package main

import "sort"

// https://leetcode.cn/problems/russian-doll-envelopes/
// 354. 俄罗斯套娃信封问题

// // 最多有多少个信封可以套娃
// // 套娃的前提是，A信封比B信封的宽和高都大才可以
// // 求最多有多少个信封能套娃，也就相当于找一个单调递增的数组，求这个数组的最大长度
// // 有点类似于 LIS，但是LIS是子序列，序列中的数，不能调换顺序，而这里可以
// // 那么还能定义dp[i] 为 以 envelopes[i] 为套娃最外信封时，套娃的长度吗
// // LIS可以是因为，它可以基于此前已经计算过的子问题，即dp[i]只考虑dp[i]前的所有dp[j],j<i，不会考虑 i 下标后的数
// // 而这里当前是全都要考虑的，如果不排序的话，就得dfs，时间复杂度就飙高了，所以不如先按第一个 w 做个升序排序
// // 这样的话，在一定程度上就可以再现 T300 的不考虑数组 i 下标后的情况了，因为 w 升序的原因，i 下标后的数对肯定不满足套娃的情况
// // 这个时候我们就可以基于 h 套用 LIS 的状态转移方程了，但是这里还有一个问题
// // LIS是单纯的看单调递增，信封同一 w 情况下，如果存在前面的h比后面的h低，那么这里显然会有问题
// // 即已排序过 w，但是 h 是无序的 [[2, 3], [5, 4], [6, 5], [6, 7]]
// // 此时基于 h 进行 LIS ---- [3,4,5,7]
// // 我们会得到长度为4的[3,4,5,7]的结果，但对于本题来说是不符合条件的，因为5,7都属于同一宽度
// // 这里有两个解决方案
// // 1. w,h 都比对，正常LIS的判断是envelopes[i][1] > envelopes[i][j], j < i // 只比对 h
// // 那么这里可以额外加一个条件 envelopes[i][1] > envelopes[j][1] && envelopes[i][0] > envelopes[j][0] , j < i
// // 即，不仅比对 h，还额外比对 w，但是这样有个缺点就是w的排序后剪枝效果没那么显著，仅仅只是排除了 i 下标后的情况
// // 2. 将h进行降序，即同一w情况下，降序h排列
// // 此时正常 LIS 即可，效果更佳比前者更佳
// // 不过这种解法也会超时(看了一些题解，发现以前不会超时，后来增加了一些测试用例就算超时了)
// func maxEnvelopes(envelopes [][]int) int {
// 	// 按照 w 升序，同 w 时 h 降序
// 	sort.Slice(envelopes, func(i, j int) bool {
// 		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
// 	})
// 	// 记录以 envelopes[i] 为结尾的最长上升子序列
// 	dp := make([]int, len(envelopes))
// 	for i := range envelopes {
// 		dp[i] = 1
// 	}
// 	res := 1
// 	for i := 1; i < len(envelopes); i++ {
// 		for j := 0; j < i; j++ {
// 			if envelopes[i][1] > envelopes[j][1] {
// 				dp[i] = max(dp[i], dp[j]+1)
// 			}
// 		}
// 		res = max(res, dp[i])
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// 那么再换 T300 的贪心+二分，还是先排序，将本题的求解过程弱化为T300的形式
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})
	// tail[i] 表示 i+1 长度下，结尾的数字最小是哪个数
	tail := make([]int, 0)
	for _, envelope := range envelopes {
		idx := findInsertIdx(tail, envelope[1])
		if idx == len(tail) {
			tail = append(tail, envelope[1])
		} else {
			tail[idx] = envelope[1]
		}
	}
	return len(tail)
}

// 在 tail 数组中找到 target 将要插入的位置
func findInsertIdx(tail []int, target int) int {
	l, r := -1, len(tail)
	for l+1 != r {
		mid := l + (r-l)/2
		if tail[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

// // 判断a能否包含b，即b信封是否可以放进a信封，true表示可以
// func canAContainB(a, b []int) bool {
// 	if a[0] <= b[0] || a[1] <= b[1] {
// 		return false
// 	}
// 	return true
// }
