package main

import (
	"sort"
)

// https://leetcode.cn/problems/merge-intervals/
// 56. 合并区间

// 首先要对区间都进行排序，不然我们不方便寻找重叠的区间
// 排序后，以第一个区间的右边界为基准，去遍历第2/3/4/5...个区间的左边界，是否在第一个区间的右边界中 <= 状态
// 满足，则将第 2/3/4.. 个区间的右边界，赋值到第一个区间的右边界，更新重叠区间的右边界长度(以大的赋值)
// 当遇到的新区间的左边界不在之前重叠区间的右边界范围内，则将这个区间作为新的重叠区间，重复上面的操作
func merge(intervals [][]int) [][]int {
	// 根据左边界先排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
			continue
		}
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
