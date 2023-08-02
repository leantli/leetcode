package main

// https://leetcode.cn/problems/next-greater-element-i/
// 496. 下一个更大元素 I

// 二刷
// 找出nums1中每个数在nums2中所在位置的下一个更大元素
// 这里我们显然先对 nums2 进行处理，求出 nums2 中每个数的下一个更大元素，放入 map 里面
// 再基于 map 和 nums1 中的数，获取对应的下一个更大元素即可
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0)
	m := make(map[int]int)
	for i := range nums2 {
		// 基于单调栈去获取每个元素的下一个更大元素
		// 栈为空时直接入数字，当栈顶的数字小于当前 nums2[i]，显然栈顶数的下一个更大元素就是 nums2[i]
		// 此时将这两个数值存入 map 中，并将栈顶弹出
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			m[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	// 这里也可以像第一次写一样，直接 map ok 判断，无则赋值为 -1
	// 此时栈中还有数，则说明该数没有下一个更大元素，将其弹出存入map，value 为 -1
	for len(stack) > 0 {
		m[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}
	// 遍历 nums1 中的每个数，根据 nums1[i] 获取 map 中的 value 即可
	ans := make([]int, len(nums1))
	for i := range nums1 {
		ans[i] = m[nums1[i]]
	}
	return ans
}

// // 和 每日温度 很像，但是这里只需要求出 nums1 中对应的数字在 nums2 中右侧的第一大数值
// // 也就是说，nums2 长度一定大于等于 nums1，我们可以先求出 nums2 中每个数的右侧的第一大数值，存入 map 中
// // 然后再遍历 nums1, 通过 map 取出对应的右侧的第一大值
// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// 	m := make(map[int]int)
// 	stack := make([]int, 0)
// 	for i := 0; i < len(nums2); i++ {
// 		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
// 			m[stack[len(stack)-1]] = nums2[i]
// 			stack = stack[:len(stack)-1]
// 		}
// 		stack = append(stack, nums2[i])
// 	}
// 	res := make([]int, len(nums1))
// 	for i := range nums1 {
// 		if num, ok := m[nums1[i]]; ok {
// 			res[i] = num
// 		} else {
// 			res[i] = -1
// 		}
// 	}
// 	return res
// }
