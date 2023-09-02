package main

// https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
// 4. 寻找两个正序数组的中位数

// 虽然成功解出来了，但是过程分类太多太杂，很累，看看官解的类似思路
// 两个数组都是正序数组，显然最简单的做法是双指针遍历
// 双指针指向两个数组的头部，每次都比较后移动教小的数组的指针
// 直到操作第(n+m)/2次(奇数) 或 (n+m)/2和(n+m)/2-1次(偶数)
// 但这样一来，时间复杂度为 O((m+2)/2) 不满足题目要求
// 那么看到 log，我们应该想到，这道题应该得用二分
// 那么这道题怎么使用二分呢？
// 二分要有确切的缩减方向，可以计算本次要排除的数量为mid
// 从两个数组中排除mid数量的数字，直到最后排除一半
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n&1 == 1 {
		return float64(findKthNum(nums1, nums2, n/2+1))
	}
	a := float64(findKthNum(nums1, nums2, n/2))
	b := float64(findKthNum(nums1, nums2, n/2+1))
	return (a + b) / 2.0
}

// 这个k是第k个，映射到下标需要减1
// 找到第k个数并会返回
func findKthNum(nums1, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	half := k/2 - 1
	// 这里用 index1 和 index2 确认要比较的下标
	index1, index2 := min(len(nums1)-1, half), min(len(nums2)-1, half)
	// 比较完，排除一部分数，排除的个数是 index1 + 1 或者 index2 + 1，index 是下标，排除个数需要 +1
	if nums1[index1] < nums2[index2] {
		return findKthNum(nums1[index1+1:], nums2, k-index1-1)
	}
	return findKthNum(nums1, nums2[index2+1:], k-index2-1)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// // 时间复杂度为 log(m+n)，大概率是使用二分的思路，但是我们又不能把两个数组拼接起来排序再二分，那样时间复杂度超了
// // 这里我们可以采用排除 k/2 个的方式去排掉两个数组中中点之前的数字，时间复杂度为 logn+logm
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	cnt := len(nums1) + len(nums2)
// 	if cnt&1 == 1 {
// 		return getK(nums1, nums2, cnt/2+1)
// 	}
// 	return (getK(nums1, nums2, cnt/2) + getK(nums1, nums2, cnt/2+1)) / 2.0
// }
// // 获取两个数组中的第 k 个数，注意这里是第 k 个，不是下标 k
// // 当 k 等于 1 或者两个数组有一个已经全部排空了，则可以直接返回对应的中点数值了
// func getK(nums1, nums2 []int, k int) float64 {
// 	if len(nums1) == 0 {
// 		return float64(nums2[k-1])
// 	}
// 	if len(nums2) == 0 {
// 		return float64(nums1[k-1])
// 	}
// 	if k == 1 {
// 		return float64(min(nums1[0], nums2[0]))
// 	}
// 	// 当 k 不为 1 时，采用 k/2=mid，排除 mid 个数字
// 	mid := k / 2
// 	// 接下来分类讨论
// 	if len(nums1) > mid && len(nums2) > mid {
// 		if nums1[mid-1] > nums2[mid-1] {
// 			return getK(nums1, nums2[mid:], k-mid)
// 		} else {
// 			return getK(nums1[mid:], nums2, k-mid)
// 		}
// 	} else if len(nums1) > mid {
// 		if nums1[mid-1] > nums2[len(nums2)-1] {
// 			return getK(nums1[mid-len(nums2):], []int{}, k-mid)
// 		} else {
// 			return getK(nums1[mid:], nums2, k-mid)
// 		}
// 	} else {
// 		if nums2[mid-1] > nums1[len(nums1)-1] {
// 			return getK([]int{}, nums2[mid-len(nums1):], k-mid)
// 		} else {
// 			return getK(nums1, nums2[mid:], k-mid)
// 		}
// 	}
// }
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
