package main

// https://leetcode.cn/problems/koko-eating-bananas/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 875. 爱吃香蕉的珂珂

// 常规，枚举 k，判断能否在 h 小时内吃完，取能吃完的最小 k
// 因为 k 过大时，消耗时间远小于 h；k 过小时，消耗时间远大于 k
// k 与 消耗时间 的关系具备单调性，因此可以采用 二分
// 当 耗时小于等于 h 时，k 都是可取的，我们只要取最小的 k 即可
func minEatingSpeed(piles []int, h int) int {
	var max int
	for _, pile := range piles {
		if max < pile {
			max = pile
		}
	}
	// // l,r 取值为[1,max(piles)]，下面采用开区间
	l, r := 0, max+1
	for l+1 != r {
		mid := l + (r-l)/2
		// 当前 mid 速度小，求耗时 hours
		var hours int
		for _, pile := range piles {
			hours += ((pile - 1) / mid) + 1
		}
		// 耗时少说明 mid 太大，可以让 mid 再小点
		if hours <= h {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}
