package main

// https://leetcode.cn/problems/maximum-distance-between-a-pair-of-values/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 1855. 下标对中的最大距离

// 虽然通过了，但是击败百分比有点低。。感觉应该是有更好的方法的
// 其实也可以单纯用双指针？
// i 先不变，j 直接遍历过去，一直到 nums2 中最后一个大于等于 nums[i]，此时计算 dis
// 然后再移动 i，j 再判断有没有必要移动，没必要的话计算 dis
// 如此往复，时间复杂度为 O(n+m)
func maxDistance(nums1 []int, nums2 []int) int {
	var i, j int
	var ans int
	for i < len(nums1) && j < len(nums2) {
		for j < len(nums2) && nums1[i] <= nums2[j] {
			j++
		}
		i++
		if i > j {
			continue
		}
		ans = max(ans, j-i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 两个数组都是非递增
// // i 是 nums1 的下标， j 是 nums2 的下标
// // 满足 i <= j && nums1[i] <= nums2[j]，则可计算下标对距离 j-i
// // 最终目的是求 最大下标对距离
// // 这样的话，我们需要对 nums1 中的每个数，去寻找一个在 nums2 的数，该数是 nums2 中最后一个大于等于 nums1 中定位的数
// // 这个寻找的过程显然使用二分是最优解，正常做法应该是 n*m 的时间复杂度，使用二分后是 n*log(m/2)
// // 先写写看
// func maxDistance(nums1 []int, nums2 []int) int {
// 	m := len(nums2)
// 	var ans int
// 	for i, num := range nums1 {
// 		// 这里还要简单特殊处理一下，因为 nums1 长度可能比 nums2 长，而 i 必定比 j 大的部分就可以直接省略了
// 		if i >= m {
// 			continue
// 		}
// 		// l 和 r 的初始化也需要简单注意一下，最后一个大于等于 nums1 的下标的位置可能范围是 [i, m-1]
// 		// 因此我们的左右指针分别是 i-1, m-1+1
// 		l, r := i-1, m
// 		for l+1 != r {
// 			mid := l + (r-l)/2
// 			if nums2[mid] >= num {
// 				l = mid
// 			} else {
// 				r = mid
// 			}
// 		}
// 		// 此时 l 会停在最后一个大于等于 num 的位置，r 会停在第一个小于 num 的位置
// 		// 当然，这里也有可能 l 没有动弹，j >= i 的情况下，没有 nums[j] > num
// 		if l == i-1 {
// 			continue
// 		}
// 		ans = max(ans, l-i)
// 	}
// 	return ans
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
