package main

import "sort"

// https://leetcode.cn/problems/non-overlapping-intervals/
// 435. 无重叠区间

// // 返回需要移除区间的最小数量，使剩余区间互不重叠
// // 一道很隐蔽的LIS变种题，区间不重叠如何判断？
// // interval 的 start 和 end 都大于另一个 interval的 start 和 end
// // 这个条件一转化出来，就有点像 二维属性的 LIS 了 (详情参考 T345 和 T17.08)
// // 当然，这里还要注意，后者的 start 还得大于等于前者的 end，并且基于这个条件，就无需判断两次大于
// // 此时求出最长的不重叠区间，再将当前区间集合的长度减去最长不重叠区间的长度
// // 就得到了需要移除区间的最小数量
// // 也就是这道题的题意相当于，选出最多数量的区间，使它们区间互不重叠，此时要移除的重叠区间数量最少
// // 不过用这种方法最后会超时，此时再接着考虑贪心
// func eraseOverlapIntervals(intervals [][]int) int {
// 	// 直接根据左端点升序即可，后续只要j的右端点小于等于i的左端点即可
// 	sort.Slice(intervals, func(i, j int) bool {
// 		return intervals[i][0] < intervals[j][0]
// 	})
// 	dp := make([]int, len(intervals))
// 	for i := range intervals {
// 		dp[i] = 1
// 	}
// 	longestLen := 1
// 	for i := 1; i < len(intervals); i++ {
// 		for j := 0; j < i; j++ {
// 			if intervals[j][1] <= intervals[i][0] {
// 				dp[i] = max(dp[i], dp[j]+1)
// 			}
// 		}
// 		longestLen = max(longestLen, dp[i])
// 	}
// 	return len(intervals) - longestLen
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// 下面考虑贪心思路，常规dp会超时
// 注意这道题LIS变种题的贪心，无需结合二分，为什么？
// 这道题的区间不重叠条件特性，完全无需二分进行额外的检索
// 根据排序以及不重叠的条件即可
// 总结两种端点排序的思路，其实本质都是一样，贪心策略是 计算最多能组成的不重叠区间个数
// 在重叠区间中，寻找位置最左的右端点，这样能保证后续非重叠区间的可选更多，更能延长非重叠区间集合的长度

// // 贪心思路 -- 左端点排序
// func eraseOverlapIntervals(intervals [][]int) int {
// 	// 根据左端点升序
// 	sort.Slice(intervals, func(i, j int) bool {
// 		return intervals[i][0] < intervals[j][0]
// 	})
// 	// tail[i] 表示 i+1 长度时，最末尾的 end 是多少？
// 	tail := []int{intervals[0][1]}
// 	for _, interval := range intervals {
// 		// intervals 已经根据左端点排序了
// 		// 当 左端点 已经大于等于 tail 数组中 end 的长度，就说明已经可以考虑下一个不重叠区间了
// 		if interval[0] >= tail[len(tail)-1] {
// 			tail = append(tail, interval[1])
// 		} else if interval[1] < tail[len(tail)-1] {
// 			// 否则的话，此时区间是重叠的，可以看tail的end是否大于当前interval的end，是的话则替换成更小的
// 			tail[len(tail)-1] = interval[1]
// 		}
// 	}
// 	return len(intervals) - len(tail)
// }

// // 贪心思路 -- 左端点排序 -- 简单优化空间
// // 注意到其实对 tail 数组的使用，我们只需要 tail 中 end 即可，无需 tail 中前面的其他 end
// // 因此可以简化为两个变量
// func eraseOverlapIntervals(intervals [][]int) int {
// 	// 根据左端点升序
// 	sort.Slice(intervals, func(i, j int) bool {
// 		return intervals[i][0] < intervals[j][0]
// 	})
// 	cnt, end := 1, intervals[0][1]
// 	for _, interval := range intervals {
// 		if interval[0] >= end {
// 			end = interval[1]
// 			cnt++
// 		} else if interval[1] < end {
// 			end = interval[1]
// 		}
// 	}
// 	return len(intervals) - cnt
// }

// 贪心思路 -- 右端点排序
func eraseOverlapIntervals(intervals [][]int) int {
	// 根据右端点升序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	// 其实我们的根本目的就是，在重叠区间中，寻找右端点最靠左的区间
	// 当下一个不重叠区间的左端点大于等于这个右端点时，就可以了延长非重叠区间的长度
	// 而此时我们已经根据右端点升序排序了，那么此时首个interval的右端点就一定是所有非重叠区间中最左区间的右端点
	// 并且接下来，当我们每次遇到左端点大于当前 end 的 interval 时，都说明进入到了一个 新的 非重叠区间
	// 并且由于右端点已经排序过了，此时新重叠区间的最左右端点，就一定是这个 interval 的右端点
	cnt, end := 1, intervals[0][1]
	for _, interval := range intervals {
		if interval[0] >= end {
			cnt++
			end = interval[1]
		}
	}
	return len(intervals) - cnt
}
