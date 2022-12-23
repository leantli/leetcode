package main

// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 51. 数组中的逆序对

// 逆序对，并且要注意是未排序后的
// 这里可以显然可以基于归并排序去统计，在并时判断左侧数组的值是否大于右侧数组的最大值
// 也就是使用归并排序，排成降序，每次左侧首位大于右侧首位，则逆序对数量 += len(右侧数组)

func reversePairs(nums []int) int {
	var res int
	var merge func(left, right []int) []int
	merge = func(left, right []int) []int {
		temp := make([]int, 0, len(left)+len(right))
		for len(left) != 0 && len(right) != 0 {
			if left[0] > right[0] {
				res += len(right)
				temp = append(temp, left[0])
				left = left[1:]
			} else {
				temp = append(temp, right[0])
				right = right[1:]
			}
		}
		temp = append(temp, left...)
		temp = append(temp, right...)
		return temp
	}
	var mergeSort func(nums []int) []int
	mergeSort = func(nums []int) []int {
		if len(nums) <= 1 {
			return nums
		}
		mid := len(nums) / 2
		left := mergeSort(nums[:mid])
		right := mergeSort(nums[mid:])
		return merge(left, right)
	}
	mergeSort(nums)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
