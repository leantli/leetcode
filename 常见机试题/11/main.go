package main

import (
	"fmt"
	"sort"
)

/**
题目描述:
给定一组闭区间，其中部分区间存在交集。任意两个给定区间的交集，称为公共区间(如: [1,21,[2,3]的公共区间为[2,21,[3,5],[3,]的公共区间为[3,5])。公共区间之间若存在交集，则需要合并(如: [1,31,[3,5]区间存在交集[3,3]，须合并为[1,5])。按升序排列输出合并后的区间列表。
输入描述:
组区间列表
区间数为N:
0 <= N <= 1000:
区间元素为X:
10000 <= X <= 10000.
输出描述:
升序排列的合并后区间列表
**/

// 先求交集，再在交集基础上合并区间
func filter(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][1] >= intervals[j][0] {
				res = append(res, []int{intervals[j][0], intervals[i][1]})
			}
		}
	}
	if len(res) == 0 {
		return [][]int{}
	}
	return merge(res)
}

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

func main() {
	fmt.Println(filter([][]int{{0, 3}, {1, 3}, {3, 5}, {3, 6}}))
	fmt.Println(filter([][]int{{0, 3}, {1, 4}, {4, 7}, {5, 8}}))
	fmt.Println(filter([][]int{{1, 2}, {3, 4}}))
}
