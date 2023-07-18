package main

// https://leetcode.cn/problems/next-greater-element-ii/
// 503. 下一个更大元素 II

func nextGreaterElements(nums []int) []int {
	// 循环数组，直接将数组 * 2 拼接，然后正常单调栈求出每个数的下一个更大元素
	// 最后截取原数组长度即可
	n := len(nums)
	nums = append(nums, nums...)
	stack := make([]int, 0)
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			preIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[preIdx] = nums[i]
		}
		stack = append(stack, i)
	}
	for _, v := range stack {
		res[v] = -1
	}
	return res[:n]
}
