package main

import "fmt"

// https://leetcode.cn/problems/median-of-two-sorted-arrays/?favorite=2cktkvj
// 4. 寻找两个正序数组的中位数

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
	index1, index2 := min(len(nums1)-1, half), min(len(nums2)-1, half)
	// 先不考虑一些特殊例子，先写出一个通用的情况
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

// // 显然需要用到二分去找到第 len/2 和 len/2 -1 位置的值
// // 但是两个数组不好继续二分找，结合题目，两个数组都是正序数组，
// // 此时转换问题---> 找到最小的第 k 个数
// // 但是两个数组，如何最小代价地找到第 k 个数？
// // 正常情况下，可能会才去遍历，但是时间复杂度显然不够，看到 log(n)，显然是要用二分
// // 那么怎么用二分好？
// // 我们也无法说先将两个数组合起来再二分，这样的话显然脱裤子放屁了
// // 也就是说，我们必须对两个数组都分别二分
// // 再考虑到两个都是正序，我们可以通过二分，先找到两个数组中二分位置的值
// // nums1[k/2] 和 nums2[k/2]，此时比较二者大小，将小的数组前 k 个值直接排去
// // 没错，就是基于 k，采用二分不断去排除两个数组前 k 个值
// // 当然，这里还有一些细节得处理，比如某个数组的 k/2 下标已经越界了等
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	n := len(nums1) + len(nums2)
// 	if n%2 == 1 {
// 		return float64(findK(nums1, nums2, n/2+1))
// 	}
// 	return float64((findK(nums1, nums2, n/2) + findK(nums1, nums2, n/2+1))) / 2.0
// }

// // 注意，这里的 k 传入的是第 k 个，而不是下标为 k
// // 因此后面一些对下标的处理，需要 k - 1 处理
// // 这道题的边界处理有点恶心
// func findK(nums1, nums2 []int, k int) int {
// 	index1, index2 := 0, 0
// 	for {
// 		if index1 == len(nums1) {
// 			return nums2[index2+k-1]
// 		}
// 		if index2 == len(nums2) {
// 			return nums1[index1+k-1]
// 		}
// 		if k == 1 {
// 			return min(nums1[index1], nums2[index2])
// 		}
// 		half := k / 2
// 		newIndex1 := min(index1+half, len(nums1)) - 1
// 		newIndex2 := min(index2+half, len(nums2)) - 1
// 		if nums1[newIndex1] <= nums2[newIndex2] {
// 			k -= (newIndex1 - index1 + 1)
// 			index1 = newIndex1 + 1
// 			continue
// 		}
// 		k -= (newIndex2 - index2 + 1)
// 		index2 = newIndex2 + 1
// 	}
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}))
}

// 马上想到归并，但显然不满足时间复杂度 O(log (m+n))，最多也只能达到 O((m+n)/2)
// 再想到根据 len/2 和 len/2 -1 位置去遍历，其实也是不行的，都是 O(m+n) 级别的
// 显然还是得靠二分
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	len1, len2 := len(nums1), len(nums2)
// 	temp := make([]int, 0, len1+len2)
// 	mid := (len1 + len2) / 2
// 	l, r := 0, 0
// 	for l < len1 && r < len2 && len(temp) <= mid {
// 		if nums1[l] < nums2[r] {
// 			temp = append(temp, nums1[l])
// 			l++
// 			continue
// 		}
// 		temp = append(temp, nums2[r])
// 		r++
// 	}
// 	temp = append(temp, nums1[l:]...)
// 	temp = append(temp, nums2[r:]...)
// 	if len(temp)%2 == 0 {
// 		return float64(temp[mid]+temp[mid-1]) / 2.0
// 	}
// 	return float64(temp[mid])
// }
