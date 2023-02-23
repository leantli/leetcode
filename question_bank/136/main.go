package main

// https://leetcode.cn/problems/single-number/
// 136. 只出现一次的数字

// 相似题型：剑指03,56-1;;leetcode-287,26,136,137,260

// 某个元素只出现了一次，其他元素都出现了两次
// 时间复杂度O(n)以下，空间复杂度O(1)
// 常规来说肯定是 map，但是不符题意；
// 要用前后比较，也得先排序，时间复杂度不满足；
// 鸠占鹊巢也不行，数组中的数字无限制
// 只能位运算
func singleNumber(nums []int) int {
	var res int
	for i := range nums {
		res ^= nums[i]
	}
	return res
}
