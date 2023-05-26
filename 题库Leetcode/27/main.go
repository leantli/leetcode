package main

// https://leetcode.cn/problems/remove-element/
// 27. 移除元素

// 只需要原地移除所有等于 val 的元素
// 原地利用数组显然需要两个指针
// 一个指针指向“新”数组当前要插入的位置
// 另一个指针遍历原数组
func removeElement(nums []int, val int) int {
	var new int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[new] = nums[i]
			new++
		}
	}
	return new
}
