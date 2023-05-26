package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/
// 80. 删除有序数组中的重复项 II

// 基于原数组维护一个新的数组，其中重复出现的元素最多可以在新数组中重复出现两次
func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	// 新数组中的最后一个数的下标
	j := 1
	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[j] {
			j++
			nums[j] = nums[i]
			continue
		}
		// 如果不满足上面的 if，那么 nums[i] 必定与 nums[j] 相等
		// 此时应判断 nums[i] 是否还和 nums[j-1] 相等，若相等，则说明新数组中已经有两个重复的该元素了，可以直接略过
		// 不等的话，就可以将整个元素放入新的数组对应的位置
		if nums[i] != nums[j-1] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
