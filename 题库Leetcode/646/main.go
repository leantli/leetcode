package main

import "sort"

// https://leetcode.cn/problems/maximum-length-of-pair-chain/
// 646. 最长数对链

// // 数对中 l < r，感觉就有点类似区间对？
// // p1[a,b], p2[c,d]，当c > b 时，能构建数对联
// // 感觉这和无重叠区间差不多了
// // 但这也和二维属性的LIS很类似，我们可以先用 LIS dp 尝试解决
// func findLongestChain(pairs [][]int) int {
// 	sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })
// 	// dp[i] 定义为 pairs[i] 作为数对结尾时最长数对链的长度
// 	dp := make([]int, len(pairs))
// 	// 初始化，每个数对都能单独成链，长度为1
// 	for i := range pairs {
// 		dp[i] = 1
// 	}
// 	res := 1
// 	for i := 1; i < len(pairs); i++ {
// 		for j := 0; j < i; j++ {
// 			if pairs[i][0] > pairs[j][1] {
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

// 再考虑贪心策略 -- 计算最大的无重叠区间是多少
// 那么我们怎么考虑贪心？
// 首先根据区间的右端点进行升序排序，此时第一个 pairs 的右端点，一定是第一个不重叠区间的最小右端点
// 当我们还没遇到左端点大于该最小右端点的区间时，此时遍历遇到的 pairs 的区间，都一定和这个第一个不重叠区间重叠
// 因为他们的左端点此时都小于等于第一个不重叠区间的右端点
// 而当我们遇到第一个左端点大于该最小右端点的区间时，这个区间，就是第二个不重叠区间
// 即，要挑选最长数对链的第一个数对时，最优的选择是挑选数对 r 最小的，这样能给挑选后续的数对留下更多的空间。
func findLongestChain(pairs [][]int) int {
	// 根据右端点升序
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	cnt, end := 1, pairs[0][1]
	for _, pair := range pairs {
		if pair[0] > end {
			end = pair[1]
			cnt++
		}
	}
	return cnt
}
