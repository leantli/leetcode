package main

import "sort"

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/
// 452. 用最少数量的箭引爆气球

// 至少要用多少箭才能引爆所有气球，显然取决于气球的区间，有多少个不重叠区间
// 因此这里我们转化为求非重叠区间的数量，顺带关注一下当 end_i == start_j 时，属于重叠
// 常规来看，我们可以基于 start_i > end_j, start_i > start_j, end_i > end_j
// 会觉得这题和 二维属性的 LIS 问题很相似，都是求最长递增子序列，可惜这道题用这类方法会超时
// 我们最终还是考虑贪心策略---计算不重叠区间有多少个
// 那么我们怎么考虑贪心？
// 首先根据气球的右端点进行升序排序，此时第一个points的右端点，一定是第一个不重叠区间的最小右端点
// 当我们还没遇到左端点大于该最小右端点的区间时，此时遍历遇到的 points 的区间，都一定和这个第一个不重叠区间重叠
// 因为他们的左端点此时都小于等于第一个不重叠区间的右端点
// 而当我们遇到第一个左端点大于该最小右端点的区间时，这个区间，就是第二个不重叠区间
func findMinArrowShots(points [][]int) int {
	// 对气球区间的右端点升序排序
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	cnt, end := 1, points[0][1]
	for _, point := range points {
		if point[0] > end {
			end = point[1]
			cnt++
		}
	}
	return cnt
}
