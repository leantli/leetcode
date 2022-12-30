package main

// https://leetcode.cn/problems/find-peak-element/
// 162. 寻找峰值

// 首先，峰值一定是大于其左右数值的
// 其次，我们要使用二分，需要找到一定的趋势
// 这个趋势，显然是峰值的左侧，必然有一段单调递增的区间
// 此时区间内单调递增
func findPeakElement(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	l, r := -1, len(nums)
	for l+1 != r {
		mid := l + (r-l)/2
		// 对特别的可能存在的越界情况做一下处理
		if mid+1 == len(nums) && nums[mid] > nums[mid-1] {
			return mid
		}
		// 常规 condition，当仍单调递增时就移动 l
		if nums[mid] < nums[mid+1] {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}
