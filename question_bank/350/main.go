package main

import "sort"

// https://leetcode.cn/problems/intersection-of-two-arrays-ii/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 350. 两个数组的交集 II

// ? 怎么分在二分计划里面，这道题用二分的话效率很低啊
// 肯定无脑哈希表啊，如果是已经排序好的话，可以考虑双指针

// 再尝试一下进阶的排序的情况
func intersect(nums1 []int, nums2 []int) []int {
	// 先排序一遍，其实后面用双指针即可，不等时，移动较小的数所在数组的指针
	// 相等时，就将该数添加到 ans，并且两个数组的指针都要移动
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
	sort.Slice(nums2, func(i, j int) bool { return nums2[i] < nums2[j] })
	var i, j int
	ans := make([]int, 0)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			ans = append(ans, nums1[i])
			i++
			j++
			continue
		}
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

// // 看完题解，再尝试一下只用一个哈希表整
// // 确实更简介更明了
// func intersect(nums1 []int, nums2 []int) []int {
// 	m := make(map[int]int)
// 	for _, num := range nums1 {
// 		m[num]++
// 	}
// 	ans := make([]int, 0)
// 	for _, num := range nums2 {
// 		if m[num] > 0 {
// 			m[num]--
// 			ans = append(ans, num)
// 		}
// 	}
// 	return ans
// }

// // 先按无排序的情况先写一下
// // 因为存在重复出现的情况，并且要按出现次数较小的数组计算，因此应分成两个哈希表进行存储
// // 比如 [1,2,2,2,2] 和 [2,2]，返回结果应是 [2,2]，而不能只是 [2]
// func intersect(nums1 []int, nums2 []int) []int {
// 	m1, m2 := make(map[int]int), make(map[int]int)
// 	for _, num := range nums1 {
// 		m1[num]++
// 	}
// 	for _, num := range nums2 {
// 		m2[num]++
// 	}
// 	ans := make([]int, 0)
// 	for k, v := range m1 {
// 		v2 := m2[k]
// 		// 取两个 map 中，出现次数少的进行添加，若为 0 也没关系
// 		count := min(v2, v)
// 		for i := 0; i < count; i++ {
// 			ans = append(ans, k)
// 		}
// 	}
// 	return ans
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
