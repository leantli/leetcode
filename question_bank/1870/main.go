package main

// https://leetcode.cn/problems/minimum-speed-to-arrive-on-time/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1870. 准时到达的列车最小时速

// 枚举最小时速，然后模拟计算该时速下到达终点耗时为 cost
// 这里枚举我们加油通过二分去找到最小时速, 设 时速 为 mid
// 模拟计算基于时速 mid 下，需要消耗的时间 cost
// 当 cost < hour 时，说明还有多余时间可以通勤，时速可以再小一点，r = mid
// 当 cost > hour 时，说明超过了通勤时间，时速得大一点，l = mid
func minSpeedOnTime(dist []int, hour float64) int {
	var max int
	for _, num := range dist {
		if num > max {
			max = num
		}
	}
	// 时速取值范围为 [1,max]，下面采用开区间进行二分
	max = max*100 + 1
	l, r := 0, max
	for l+1 != r {
		mid := l + (r-l)/2
		cost := getCost(mid, dist)
		if cost <= hour {
			r = mid
		} else {
			l = mid
		}
	}
	if r == max {
		return -1
	}
	return r
}

// 根据时速 mid 和 路程，计算返回全路程耗时
func getCost(mid int, dist []int) float64 {
	var cost int
	// 对于 dist 除了最后一个路程外，其他路程除以时速后都得向上取整
	pre := len(dist) - 1
	for i := 0; i < pre; i++ {
		cost += (dist[i]-1)/mid + 1
	}
	return float64(cost) + float64(dist[len(dist)-1])/float64(mid)
}
