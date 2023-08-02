package main

// https://leetcode.cn/problems/next-greater-element-ii/
// 503. 下一个更大元素 II

// 二刷
// 下一个更大元素的进阶题，显然，我们还是需要利用单调栈去获取当前元素的下一个更大元素
// 但这次包含了循环的条件，我们可以通过数组*2，简单解决循环导致的下一个最大元素在左侧的case
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nums = append(nums, nums...)
	ans := make([]int, len(nums))
	stack := make([]int, 0)
	for i := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			ans[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		// 栈中存下标，确定该下标后第一个更大元素
		// 为什么不像 496 直接存对应的数值？因为 496 中元素唯一，而这里元素可重复
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		ans[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}
	return ans[:n]
}

// func nextGreaterElements(nums []int) []int {
// 	// 循环数组，直接将数组 * 2 拼接，然后正常单调栈求出每个数的下一个更大元素
// 	// 最后截取原数组长度即可
// 	n := len(nums)
// 	nums = append(nums, nums...)
// 	stack := make([]int, 0)
// 	res := make([]int, len(nums))
// 	for i := 0; i < len(nums); i++ {
// 		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
// 			preIdx := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			res[preIdx] = nums[i]
// 		}
// 		stack = append(stack, i)
// 	}
// 	for _, v := range stack {
// 		res[v] = -1
// 	}
// 	return res[:n]
// }
