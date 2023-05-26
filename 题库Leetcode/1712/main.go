package main

// https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1712. 将数组分成三个子数组的方案数

// 求 好的 分割方案数目
// 什么算 好的
// 1. 数组被分成 三个 非空连续子数组
// 2. 左数组和 <= 中数组和 <= 右数组和

// 先来个前缀和？此时复杂度为 O(n)
// 由于是三个连续非空子数组
// 因此左数组(第一个子数组)的右边界，可以直接依靠 遍历 + 其前缀和必须小于等于右边数值总和的 1/2 这个条件去确定
// 再对中数组和右数组基于前缀和求出 第一个中数组和 <= 右数组和 以及 最后一个 中数组和 <= 右数组和 的边界，相减得到左数组固定情况下，后二者的可能数目
// 时间复杂度为 O(n*logn)
func waysToSplit(nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, num := range nums {
		sum[i+1] = sum[i] + num
	}
	var ans int
	for i := 0; i < n-2; i++ {
		if sum[i+1] > (sum[n]-sum[i+1])/2 {
			continue
		}
		// 要先保证中数组和大于等于左数组
		// 接下来先用二分找到 中数组和大于等于左数组和的 第一个下标
		l, r := i, n
		for l+1 != r {
			mid := l + (r-l)/2
			if sum[mid+1]-sum[i+1] < sum[i+1] {
				l = mid
			} else {
				r = mid
			}
		}
		// l 落在最后一个不满足 好的 的下标，r 落在第一个 使得中数组和大于等于左数组和 下标
		pre := r
		// 接下来再一次二分，去找到 右数组和大于等于 中数组和 的最后一个下标
		l, r = pre-1, n-1
		for l+1 != r {
			mid := l + (r-l)/2
			if sum[mid+1]-sum[i+1] <= sum[n]-sum[mid+1] {
				l = mid
			} else {
				r = mid
			}
		}
		// l 落在最后一个 满足 右数组和 大于等于 中数组和 的下标
		ans += (l - pre + 1) % (1e9 + 7)
	}
	return ans % (1e9 + 7)
}

// 总结一下，本质就是求 中数组 的最小取值和最大取值
