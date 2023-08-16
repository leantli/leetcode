package main

import (
	"sort"
)

// https://leetcode.cn/problems/merge-intervals/
// 56. 合并区间

// 合并所有重叠区间，首先我们可以依照区间的左边界进行排序
// 方便我们进行合并操作，如何合并呢？排列后呢，当后面的区间的左边界
// 还在前一个区间的右边界范围内，我们就柠檬将其合并起来，更新最新的合并后的区间的右边界大小
func merge(intervals [][]int) [][]int {
	// 依照每个区间的左边界升序排列
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	// 直接将合并区间加入到 res 数组中，其中 res[len(res)-1] 即当前正在合并的区间
	// 通过 判断 intervals[i][0] 是否小于等于 res[len(res)-1][1] 就知道是否可以将当前区间合并到 res 中最后一个合并区间中
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		// 当前区间的左边界，还在正在合并区间的右边界范围内，则更新最新的最大的右边界
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
			continue
		}
		// 当前区间的左边界在正在合并区间的右边界范围外，则说明合并的区间已经完成了，没有其他重叠区间需要合并了
		// 此时更新合并区间的右边界，再将合并区间加入 res 数组；再重新更新合并区间的左右边界
		res = append(res, intervals[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 首先要对区间都进行排序，不然我们不方便寻找重叠的区间
// // 排序后，以第一个区间的右边界为基准，去遍历第2/3/4/5...个区间的左边界，是否在第一个区间的右边界中 <= 状态
// // 满足，则将第 2/3/4.. 个区间的右边界，赋值到第一个区间的右边界，更新重叠区间的右边界长度(以大的赋值)
// // 当遇到的新区间的左边界不在之前重叠区间的右边界范围内，则将这个区间作为新的重叠区间，重复上面的操作
// func merge(intervals [][]int) [][]int {
// 	// 根据左边界先排序
// 	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
// 	res := make([][]int, 0)
// 	res = append(res, intervals[0])
// 	for i := 1; i < len(intervals); i++ {
// 		if intervals[i][0] <= res[len(res)-1][1] {
// 			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
// 			continue
// 		}
// 		res = append(res, intervals[i])
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
