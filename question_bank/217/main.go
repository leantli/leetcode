package main

// https://leetcode.cn/problems/contains-duplicate/
// 217. 存在重复元素

// 相似题型：剑指03,56-1;;leetcode-287,26,136,137,260

// 最简单的就是借用 map 了
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}
