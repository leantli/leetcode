package main

// https://leetcode.cn/problems/minimum-number-of-days-to-make-m-bouquets/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1482. 制作 m 束花所需的最少天数

// 枚举天数 day，判断 day 天后，获得的花是否大于等于 m，是则可取值，但这里我们要取最少等待的天数
// 这里枚举可以换成二分 <--> day 和 可取得的花 具备单调性，因此可以缩进
// 设 day 为 mid
// 当可制作的花束太多时，时间可缩短，r = mid
// 当可制作的花束太少，说明需要更多时间 l = mid
// 制作 m 束花，每束花需要用到相邻的 k 朵花
func minDays(bloomDay []int, m int, k int) int {
	// 花不够直接返回
	if len(bloomDay) < k*m {
		return -1
	}
	// 天数的取值范围为[1, max(bloomDay)]，下面采用开区间二分
	var max int
	for _, day := range bloomDay {
		if day > max {
			max = day
		}
	}
	l, r := 0, max+1
	for l+1 != r {
		mid := l + (r-l)/2
		cnt := getCnt(bloomDay, m, k, mid)
		if cnt >= m {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

// 获取能制作成功的花束
func getCnt(bloomDay []int, m, k, d int) int {
	var cnt int
	var temp int
	for _, day := range bloomDay {
		if day > d {
			temp = 0
			continue
		}
		temp++
		if temp == k {
			cnt++
			temp = 0
		}
	}
	return cnt
}
