package main

import "sort"

// https://leetcode.cn/problems/range-sum-of-sorted-subarray-sums/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1508. 子数组和排序后的区间和

// 1. 先计算所有非空连续子数组和，再升序排序，此时得到一个新数组
// 2. 再返回下标 [left-1,right-1] 的数字和
// 做完差不多是击败 50%，感觉是比较常规的做法，看了还能结合二分，是有点没想到
func rangeSum(nums []int, n int, left int, right int) int {
	// 基于前缀和再去算非空连续子数组的和吧，不然重复计算可能更多
	sum := make([]int, n+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}
	nNums := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			nNums = append(nNums, sum[j]-sum[i])
		}
	}
	sort.Ints(nNums)
	var ans int
	for i := left - 1; i < right; i++ {
		ans = (nNums[i] + ans) % (1e9 + 7)
	}
	return ans
}

// 二分的做法看了一遍题解 https://leetcode.cn/problems/range-sum-of-sorted-subarray-sums/solution/-by-mkdir700-l90o/
// 怎么说呢，感觉不太能想得到，即便我做过甚至看过题解，在一定时间后仍会忘记。。。
