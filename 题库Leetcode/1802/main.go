package main

// https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=bv79h7h
// 1802. 有界数组中指定下标处的最大值

// 显然数组最终形成的图像就是， index 为山峰，整个山的走势隔一个下标，abs(h) <= 1
// 这里我们基于二分，不断去确定 index 位置的值，判断其导致的 sum 是否会超出 maxSum
// 超出则 r = mid, 保证下次二分得到的 sum 缩小
func maxValue(n int, index int, maxSum int) int {
	l, r := maxSum/n-1, maxSum+1
	for l+1 != r {
		mid := l + (r-l)/2
		// 此时假设 mid 为 index 位置的值，则其两侧的值都以减一的趋势不断降低
		// 直至两侧达到边界或值降到 1 无可再降，则两侧继续都为 1
		// 这里还未持续为 1 时，其能够等差求和
		if calSum(n, index, mid) <= maxSum {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

// 根据数组的长度，最大值及其下标，计算该数组和
// 假设 6 为最大值
// 1 1 1 1 2 3 4 5 6 5  4  3  2  1 1
// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
// 3 4 5 6 5 4 3 2
func calSum(n, index, max int) int {
	var sum int
	l, r := index-max, index+max
	if l >= 0 {
		sum += (1 + max) * max / 2
		sum += l + 1
	} else {
		sum += (max - index + max) * (index + 1) / 2
	}
	if r <= n {
		sum += (1 + max) * max / 2
		sum += n - r
	} else {
		sum += (max - (n - index - 1) + max) * (n - index) / 2
	}
	return sum - max
}
