package main

import (
	"sort"
)

// https://leetcode.cn/problems/find-k-th-smallest-pair-distance/
// 719. 找出第 K 小的数对距离

// 求第 k 小的绝对差值
// 所有数都是非负数，求最小绝对差值数对，那就直接先排序一次？更方便求较小的绝对差值
// 排序一次之后，绝对差值较近的都相邻在一起，但是又如何找到第 k 小？
// 毕竟我们每次比较也只能比较两个，有一种思路是 维护一个 k 容量的大顶堆，n^n 遍历数组计算绝对差值
// 但是这样算来时间复杂度也有 n^2 了
// 这里我们注意到，求一个存在范围的整数，第 k 小的数对距离，这个距离是不确定的
// 我们可以通过二分去确定这个距离，我们基于二分，每次选取一个 数值 作为目标数对距离
// 每次二分确定好 目标数对距离后，计算 比这个目标数对距离小的 数对 有多少？
// 当数对数量大于 k，那么显然我们设置的 目标数值太大了，数对太多了，因此可以 r = mid
// 当数对数量小于 k, 那么我们设置的目标数值太小了，数对太少了，因此可以 l = mid
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	// 开区间二分，绝对差值的取值范围为[0,t-d]
	l, r := -1, nums[len(nums)-1]-nums[0]+1
	for l+1 != r {
		// mid -- 窗口中最大绝对差值小于等于 mid -- 窗口性质
		mid := l + (r-l)/2
		// -----分割下面滑窗求 cnt 部分，下面可以单独封装成一个函数，看起来会更容易理解
		var left, right, cnt int // 此处为滑窗双指针, cnt为 最大绝对差值小于等于 mid 的对数数量
		for right = left + 1; right < len(nums); right++ {
			// 当不满足窗口性质时
			for left <= right && nums[right]-nums[left] > mid {
				left++
			}
			cnt += right - left
		}
		// ------分隔上面滑窗求cnt部分
		if cnt < k {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}
