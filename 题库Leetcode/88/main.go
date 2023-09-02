package main

// https://leetcode.cn/problems/merge-sorted-array/description/
// 88. 合并两个有序数组

// 感觉没啥好说的，nums1 甚至都直接留出了空位，我们直接从 n m 开始去比较两边的数大小
// 从 len(nums1)-1 下标开始填充 nums1
func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	for i := len(nums1) - 1; i >= 0; i-- {
		if m < 0 {
			nums1[i] = nums2[n]
			n--
			continue
		}
		if n < 0 {
			nums1[i] = nums1[m]
			m--
			continue
		}
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
	}
}
