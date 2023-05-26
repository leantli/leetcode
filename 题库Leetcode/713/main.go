package main

// https://leetcode.cn/problems/subarray-product-less-than-k/
// 713. 乘积小于 K 的子数组

// 连续子数组 子数组内所有元素乘积 小于 k
// 然后是求满足条件的连续子数组的数目
// 窗口性质-窗口内所有元素乘积小于 k
func numSubarrayProductLessThanK(nums []int, k int) int {
	var l, r, cnt int
	multiRes := 1
	for r < len(nums) {
		// 正常窗口右边界右移
		multiRes *= nums[r]
		// 窗口不满足性质时维护
		for l <= r && multiRes >= k {
			multiRes /= nums[l]
			l++
		}
		r++
		cnt += r - l
	}
	return cnt
}
