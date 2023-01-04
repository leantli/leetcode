package main

// https://leetcode.cn/problems/search-insert-position/
// 35. 搜索插入位置

// 注意哈，排序数组+目标值，这不铁铁二分？
// 该模板下, l 会指向满足条件的左区域的最大值， r 会指向不满足条件的右区域的最小值
// 当然，我们需要对结果做一些后处理，比如 target 不存在时，l 需要返回将要插入的位置
// 以及 l 未动的情况
// func searchInsert(nums []int, target int) int {
// 	l, r := -1, len(nums)
// 	for l+1 != r {
// 		mid := l + (r-l)/2
// 		if nums[mid] <= target {
// 			l = mid
// 		} else {
// 			r = mid
// 		}
// 	}
// 	if l == -1 {
// 		return 0
// 	}
// 	if nums[l] != target {
// 		return l + 1
// 	}
// 	return l
// }

// 再小改条件后，返回 r ，可以省略到更多无用的后处理
func searchInsert(nums []int, target int) int {
	l, r := -1, len(nums)
	for l+1 != r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

// // 搜索插入位置，类似于准确值模板，但未搜到时，r 会落在 最大的 小于 target 的值的下标，因此此时插入位置为 r + 1
// // 未搜到时 l 落在 最小的 大于 target 的值的下标
// func searchInsert(nums []int, target int) int {
//     l, r := 0, len(nums)-1
//     for l <= r {
//         mid := l + (r-l)/2
//         if nums[mid] < target {
//             l = mid + 1
//         } else if nums[mid] > target {
//             r = mid - 1
//         } else {
//             return mid
//         }
//     }
//     return r+1
// }
