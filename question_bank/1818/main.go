package main

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/minimum-absolute-sum-difference/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1818. 绝对差值和

// 反正也排序了，时间复杂度到了 nlogn，事实上对每个 nums2[i] 来一遍二分，时间复杂度也是相同的
// 直接在逐个计算 nums1[i] 和 nums2[i] 绝对差值时，就去 二分 查找和 nums2[i] 最接近的 nums1[i]，并计算绝对差值，寻找差值小了最多的？
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	var sum int
	n := len(nums1)
	mod := 1000000007
	sorted := make([]int, n)
	copy(sorted, nums1)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	var maxDiff int
	for i := 0; i < n; i++ {
		abs := abs(nums1[i], nums2[i])
		sum += abs
		minAbs := binarySearchMinAbs(sorted, nums2[i])
		maxDiff = max(maxDiff, abs-minAbs)
	}
	return (sum - maxDiff) % mod
	// 如果是 in32, 则 (sum-maxDiff+mod)%mod
}

// 求最小绝对差值，基于 nums2[i] 二分搜索 nums1 中最接近的两个数，并分别求绝对差值，取最小绝对差值
func binarySearchMinAbs(sorted []int, num int) (minAbs int) {
	l, r := -1, len(sorted)
	for l+1 != r {
		mid := l + (r-l)/2
		if sorted[mid] <= num {
			l = mid
		} else {
			r = mid
		}
	}
	absL, absR := math.MaxInt, math.MaxInt
	if l > -1 {
		absL = abs(sorted[l], num)
	}
	if r < len(sorted) {
		absR = abs(sorted[r], num)
	}
	return min(absL, absR)
}

func abs(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// // 找到最大绝对差值和，再去改 对应位置的 nums1[i]，基于 nums2[i] 去找 nums1 中最接近的数(这里可以对 nums1 排序后二分？)，这样得到的数是否一定会得到最小差值和？
// // 先试试吧
// // 试了一下，确实是失败了 --> 存在 [1 28 21] 和 [9 21 20] 这种，虽然 1 和 9 差值最大，但是 28 和 21 这一对，将 28 替换成 21 后，总差值和会更小
// func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
// 	var sum int
// 	n := len(nums1)
// 	mod := 1000000007
// 	var maxAbs int
// 	var maxIndex int
// 	for i := 0; i < n; i++ {
// 		abs := abs(nums1[i], nums2[i])
// 		if abs > maxAbs {
// 			maxAbs = abs
// 			maxIndex = i
// 		}
// 		sum = (sum + abs) % mod
// 	}
// 	sum = (sum - maxAbs) % mod
// 	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
// 	l, r := -1, n
// 	for l+1 != r {
// 		mid := l + (r-l)/2
// 		if nums1[mid] <= nums2[maxIndex] {
// 			l = mid
// 		} else {
// 			r = mid
// 		}
// 	}
// 	absL, absR := math.MaxInt, math.MaxInt
// 	if l > -1 {
// 		absL = abs(nums1[l], nums2[maxIndex])
// 	}
// 	if r < n {
// 		absR = abs(nums1[r], nums2[maxIndex])
// 	}
// 	return (sum + min(absL, absR)) % mod
// }

// func abs(a, b int) int {
// 	res := a - b
// 	if res < 0 {
// 		return -res
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
