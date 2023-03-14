package main

// https://leetcode.cn/problems/intersection-of-two-arrays/
// 349. 两个数组的交集

// 重复出现的数也只记录一次
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	for _, v := range nums1 {
		set[v] = struct{}{}
	}
	res := make([]int, 0)
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
