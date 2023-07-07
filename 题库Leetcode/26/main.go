package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
// 26. 删除有序数组中的重复项

// 相似题型：剑指03,56-1;;leetcode-287,26,136,137,260

// 每个数组只出现一次，二刷
// 从后往前更好写？
// func removeDuplicates(nums []int) int {
// 	n := len(nums)
// 	if n < 2 {
// 		return n
// 	}
// 	for i := n - 1; i >= 1; i-- {
// 		if nums[i] == nums[i-1] {
// 			nums = append(nums[:i], nums[i+1:]...)
// 		}
// 	}
// 	return len(nums)
// }

// // 题目只需要返回删除后的长度，因此我们并不需要真的做删除操作
// // 但是该题会校验原数组，因此我们不得不对原数组进行操作
// // 有点类似与选择排序或填坑的感觉？但是只是不考虑相等的数字
// // 基于原数组维护一个新的数组
// func removeDuplicates(nums []int) int {
// 	if len(nums) < 2 {
// 		return len(nums)
// 	}
// 	var j int
// 	for i := 1; i < len(nums); i++ {
// 		// 这里判断要注意，j 始终指向满足条件的最后一个位置，因此我们要 j++ 后才赋值
// 		if nums[i] > nums[j] {
// 			j++
// 			nums[j] = nums[i]
// 		}
// 	}
// 	return j + 1
// }
