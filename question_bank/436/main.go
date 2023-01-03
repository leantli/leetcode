package main

import "sort"

// https://leetcode.cn/problems/find-right-interval/
// 436. 寻找右区间

// 暴力，对每一个区间的 end，遍历所有区间的 start，找到最小的 大于等于 end 的 start 的下标, 此时时间复杂度为 O(n^2)
// 这里可以简单优化一下，不暴力遍历，先排序一下，然后二分, 排序 nlogn,每个都二分，也是 nlogn, 总体时间复杂度仍为 O(nlogn)
func findRightInterval(intervals [][]int) []int {
	// 但是我们最后仍要返回排序前的位置，因此我们把排序前的位置前 append 到 intervals 的 下标 2 位置
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	// 然后再根据每个区间的 start 下标排序，小的在前
	// 为什么要根据 start 排？因为我们后续时要找到 最小的 大于等于 end 的 start 的区间，主要是找 start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([]int, len(intervals))
	// 排序后再对每个区间的 end 去二分查找，找到最小的 >= end 的 start
	for _, interval := range intervals {
		l, r := -1, len(intervals)
		for l+1 != r {
			mid := l + (r-l)/2
			if intervals[mid][0] < interval[1] {
				l = mid
			} else {
				r = mid
			}
		}
		// 没找到，没找到时，r 仍在原位
		// 注意，下面 ans 的下标也必须对应排序前的数的位置
		if r == len(intervals) {
			ans[interval[2]] = -1
			continue
		}
		// 此时每个 区间 的 下标 2 位置都存储着其排序前的下标
		ans[interval[2]] = intervals[r][2]
	}
	return ans
}

// // 感觉这个比较好写？
// // 还有一种双指针思路，创建两个数组，一个存储 start ，一个存储 end，并且都是排序后为升序的状态
// // 如此一来，此时以 end 为主体遍历，遇到第一个大于它的 start 即可存储对应的下标，并且根据升序的特性，两个数组各自维护一个 指针 ，各自一次遍历即可
// func findRightInterval(intervals [][]int) []int {
// 	type pair struct {
// 		x int
// 		i int
// 	}
// 	starts := make([]pair, 0, len(intervals))
// 	ends := make([]pair, 0, len(intervals))
// 	for i, interval := range intervals {
// 		starts = append(starts, pair{x: interval[0], i: i})
// 		ends = append(ends, pair{x: interval[1], i: i})
// 	}
// 	sort.Slice(starts, func(i, j int) bool { return starts[i].x < starts[j].x })
// 	sort.Slice(ends, func(i, j int) bool { return ends[i].x < ends[j].x })

// 	ans := make([]int, len(intervals))
// 	si, ei := 0, 0
// 	// 找到第一个 >= ends[ei] 的 start[si]
// 	for si < len(intervals) {
// 		for ei < len(intervals) && ends[ei].x <= starts[si].x {
// 			ans[ends[ei].i] = starts[si].i
// 			ei++
// 		}
// 		si++
// 	}
// 	for ei < len(intervals) {
// 		ans[ends[ei].i] = -1
// 		ei++
// 	}
// 	return ans
// }
