package main

// https://leetcode.cn/problems/first-missing-positive/description/
// 41. 缺失的第一个正数

// 如果允许O(n)的额外空间，那么 map 会是一个很简单的解决方案，但是这里只允许使用常数级别的额外空间
// 如果允许O(nlogn)的时间复杂度，那么排序后遍历也会是一个简单的解决方案，这里都不行
// 先考虑一个点，数组长度为 n，那么缺失的数一定是 [1,n]
// 两轮遍历，第一轮遍历，将在[1,n]位置的数都丢到对应的位置上，不在[1,n]位置的数以及 nums[x-1]=x 则不变
// 第二轮遍历遇到的第一个 nums[i] != i+1 的数，则缺失的数为 i+1，如果遍历完了没遇到，则确实的数为 n+1
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		cur := nums[i]
		for cur >= 1 && cur <= n && nums[cur-1] != cur {
			nums[i], nums[cur-1] = nums[cur-1], nums[i]
			cur = nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
