package main

import "sort"

// https://leetcode.cn/problems/next-permutation/description/
// 31. 下一个排列

// [1,2,3]->[1,3,2]->[2,1,3]->[2,3,1]->[3,1,2]->[3,2,1]->[1,2,3]
// [1,2,3,4]->[1,2,4,3]->[1,3,2,4]->[1,3,4,2]->[1,4,2,3]->[1,4,3,2]
func nextPermutation(nums []int) {
	// 从后往前选定一个位置，看这个位置是否小于后面的某个数(这个数也从后往前遍历)，是则置换，并且对于当前 idx 后面的数进行一次升序排序
	// 当遍历到 nums[0] 时，后面如果没有数大于 nums[0]，则对全部进行一次升序排序(回到字典序最小的排列)
	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				sort.Ints(nums[i+1:])
				return
			}
		}
	}
	sort.Ints(nums)
	return
}
