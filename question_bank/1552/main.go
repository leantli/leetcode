package main

import (
	"sort"
)

// https://leetcode.cn/problems/magnetic-force-between-two-balls/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1552. 两球之间的磁力

// 使任意两球间的最小磁力最大
// 去枚举 最小磁力 为 x 是，position 能放下的小球数量是否大于等于 m，能就可以取该 x
// 在能放下小球的情况下，取最大 x
// 这里我们不枚举，采用二分去判断
// 假设两球间最小磁力为 mid，整个 position 能够放下 m 个小球
// 能放下，则说明 mid 太小，还可以增大，l = mid
// 放不下，则说明 mid 太大，只能缩小，r = mid
func maxDistance(position []int, m int) int {
	sort.Slice(position, func(i, j int) bool { return position[i] < position[j] })
	// l,r 取值范围为 [1, max-min]，下面采用开区间
	l, r := 0, position[len(position)-1]-position[0]+1
	for l+1 != r {
		// 假设两球间最小磁力为 mid，整个 position 能够放下 m 个小球
		// 能放下，则说明 mid 太小，还可以增大，l = mid
		// 放不下，则说明 mid 太大，只能缩小，r = mid
		mid := l + (r-l)/2
		if check(position, mid, m) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

// 检查是否 position 篮子，能否基于 mid 磁力放下 m 个小球
func check(position []int, mid, m int) bool {
	pre, count := position[0], 1
	for i := 1; i < len(position); i++ {
		if position[i]-pre >= mid {
			pre = position[i]
			count++
		}
	}
	return count >= m
}
