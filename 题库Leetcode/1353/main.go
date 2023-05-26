package main

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/
// 1353. 最多可以参加的会议数目

// 这道题我们很容易想到 435. 无重叠区间
// 但是我们还要额外关注一个信息：在会议区间中的任意一天参加会议，其实都是可以的
// 那我们还能想到一种贪心思路
// 当前到了哪一天，我们就选择这一天能开的会里面，结束时间最早的一天
// 因为结束时间晚的会，可选择的时间更多，之后还能开的机会更多，能让我们参加更多的会
func maxEvents(events [][]int) int {
	// 对会议的开始时间进行排序，开始时间升序，同开始时间时，结束时间升序
	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] })
	curday := 1
	cnt := 0
	h := hp{}
	for len(events) > 0 || h.Len() > 0 {
		// 把当天curday开始的会议都加入小顶堆中，准备做计划安排
		for len(events) > 0 && events[0][0] == curday {
			h.push(events[0][1])
			events = events[1:]
		}
		// 堆中存的都是结束时间，当堆顶的会议都超时了，就弹出，不参加咯
		for h.Len() > 0 && h.top() < curday {
			h.pop()
		}
		// 当堆中还有计划会议，就参加这个会议
		// 这里不用判断curday是否小于等于会议结束时间
		// 因为经过了前面两个for，此时堆中的会议一定是能参加的
		if h.Len() > 0 {
			cnt++
			h.pop()
		}
		curday++
	}
	return cnt
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(x int) { heap.Push(h, x) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }
func (h *hp) top() int   { return h.IntSlice[0] }

// // 最朴素思想---哪个结束的早去哪个
// // 因此先按会议结束时间排序
// // 接着遍历他们开会的时间区间，看那天是否没有被占用
// // 不过显然会超时
// func maxEvents(events [][]int) int {
// 	// 根据结束时间升序
// 	sort.Slice(events, func(i, j int) bool { return events[i][1] < events[j][1] })
// 	// 用于存储已被占用的时间
// 	set := make(map[int]struct{})
// 	for _, event := range events {
// 		for i := event[0]; i <= event[1]; i++ {
// 			if _, ok := set[i]; !ok {
// 				set[i] = struct{}{}
// 				break
// 			}
// 		}
// 	}
// 	return len(set)
// }

// // 佬的做法
// func maxEvents(events [][]int) int {
// 	if len(events) <= 1 {
// 		return len(events)
// 	}
// 	// 首先对原本数组进行右端点进行升序排列，因为升序之后，越排在前面的结束的越早，可选择参加的天数也就少(比如是1，第一天就结束了， 第一天不参加就错过了)
// 	sort.Slice(events, func(i, j int) bool {
// 		return events[i][1] < events[j][1]
// 	})

// 	// 先制作一个并查集所需要的数组
// 	// 因为要求尽可能参加会议，所以最理想的情况下，每天都参加一个得出的结果是最大的
// 	// 辅助集合判断这一天是否参加过, 长度为结束那天的后一天。如果已经到结束的后一天说明也结束了，因为这一天不存在
// 	worked := make([]int, events[len(events)-1][1]+2)
// 	// 代表着当前天是否被占用过，如果被占用过，指向下一个没有被占用的天的位置
// 	for k := range worked {
// 		worked[k] = k
// 	}
// 	// 要返回的结果--参加会议的次数
// 	cnt := 0
// 	for k := 0; k < len(events); k++ {
// 		// 因为在找的过程中，使用的双层for循环(for start to end)，会超出时间限制，故需要优化----并查集
// 		// 从 startDay 开始寻找下一个有空的日子
// 		today := find(worked, events[k][0])
// 		// 下一个有空的日子在 endDay 范围内则可以参加
// 		if today <= events[k][1] {
// 			cnt++
// 			union(worked, today, today+1)
// 		}
// 	}
// 	return cnt
// }

// func find(parents []int, i int) int {
// 	if parents[i] != i {
// 		parents[i] = find(parents, parents[i])
// 	}
// 	return parents[i]
// }

// func union(parents []int, i, j int) {
// 	rootX, rootY := find(parents, i), find(parents, j)
// 	if rootX == rootY {
// 		return
// 	}
// 	parents[rootX] = rootY
// }
