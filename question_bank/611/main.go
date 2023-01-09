package main

import "sort"

// https://leetcode.cn/problems/valid-triangle-number/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 611. 有效三角形的个数

// 但是不得不说，排序+二分的时间复杂度还是太高了感觉
// 感觉这里的二分又可以使用滑动窗口？
// 固定 a 后，在后面的区间基于排序后滑动窗口，时间复杂度为 2n
// 总体时间复杂度为 n*2n
func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	var count int
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}
		// 明确滑动窗口的起始点
		l, r := i+1, i+2
		for l < n {
			// 当满足条件时 r 一直右扩
			for r < n && nums[i]+nums[l] > nums[r] {
				r++
			}
			// 不满足条件时 l 右移
			l++
			// 当 r 未满足过一次条件时，直接跳过
			if r == i+2 {
				continue
			}
			count += r - l
		}
	}
	return count
}

// // nums 只包含非负整数，返回可以组成三角形的三元组个数
// // 什么情况下可以组成三角形？三边中任意两边之和大于第三边即可
// // 这里我们就要找三个数，分别为 a,b,c 其满足 a + b > c，即可组成三角形
// // 问题转换后觉得类似三数之和？
// // 先用二分试试水，先确定 a 和 b，然后寻找最后一个满足小于 a + b 的数(此时该 a，b 情况下，可以直接计算出可以与 a，b 组成三角形的 c 的数量)
// // 这里也要注意，这里使用二分的前提是先排序过一次，这样才方便找到最后一个小于 a + b 的数
// // 整个时间复杂度是 排序(nlogn)+n*n*logn
// func triangleNumber(nums []int) int {
// 	// 这里要注意，非负整数，存在 0 的情况，直接跳过
// 	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
// 	n := len(nums)
// 	var count int
// 	for a := 0; a < n; a++ {
// 		if nums[a] == 0 {
// 			continue
// 		}
// 		for b := a + 1; b < n; b++ {
// 			pivot := nums[a] + nums[b]
// 			// 接下来使用二分，先确定一下二分的边界 [b+1, n-1]
// 			l, r := b, n
// 			for l+1 != r {
// 				mid := l + (r-l)/2
// 				if nums[mid] < pivot {
// 					l = mid
// 				} else {
// 					r = mid
// 				}
// 			}
// 			// 二分结束后，l 会落在最后一个小于 a+b 的位置，r 会落在第一个大于等于 a+b 的位置
// 			// 这里我们只需要关注 l 即可，也会存在 l 的位置没动，不存在 c 小于 a+b
// 			if l == b {
// 				continue
// 			}
// 			count += l - b
// 		}
// 	}
// 	return count
// }
