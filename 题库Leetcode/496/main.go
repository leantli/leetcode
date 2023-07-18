package main

// https://leetcode.cn/problems/next-greater-element-i/
// 496. 下一个更大元素 I

// 和 每日温度 很像，但是这里只需要求出 nums1 中对应的数字在 nums2 中右侧的第一大数值
// 也就是说，nums2 长度一定大于等于 nums1，我们可以先求出 nums2 中每个数的右侧的第一大数值，存入 map 中
// 然后再遍历 nums1, 通过 map 取出对应的右侧的第一大值
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	stack := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			m[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	res := make([]int, len(nums1))
	for i := range nums1 {
		if num, ok := m[nums1[i]]; ok {
			res[i] = num
		} else {
			res[i] = -1
		}
	}
	return res
}
