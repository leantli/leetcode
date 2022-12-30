package main

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
// 34. 在排序数组中查找元素的第一个和最后一个位置

// 非递减顺序排列，有目标值，找到该目标值的开始和结束位置
// 不存在则返回 [-1,-1]
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := -1, len(nums)
	for l+1 != r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}
	// 此时 r 为 target 或 大于 target 的值，属于起始位置
	if r == len(nums) {
		return []int{-1, -1}
	}
	if nums[r] != target {
		return []int{-1, -1}
	}
	res := make([]int, 0, 2)
	res = append(res, r)
	for r < len(nums) && nums[r] == target {
		r++
	}
	return append(res, r-1)
}
