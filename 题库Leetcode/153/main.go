package main

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 153. 寻找旋转排序数组中的最小值

// 二刷，不管怎么旋转，一般总是有两个升序区间，也有可能只有一个
// 但是我们基于 mid 和 r 指向的数进行比较，只要 nums[mid] < nums[r]
// 就说明，mid 这个值，是我们要选取的可能的值
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] <= nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}

// 可以观察到，旋转一次就是把最后的数提到最前面
// 并且旋转次数不定，因此该数组旋转后，可能还是原数组，全局升序
// 常规旋转后，应该是存在两个局部升序区间
// 但显然，如果存在两个局部升序区间，我们也无需搭理前面的升序区间，只用管后面那个，
// 因为我们要找的最小元素，必定在后一个升序区间的首位
// 我们就先考虑，怎么在一个升序区间中用二分找到最小值(假设不直接定位[0]的话)
// 使用模糊值模板
// 但是怎么判断缩进方向及其正确性呢？
// 里面是两个递增区间，如果 mid 大于 n-1，
// 那必定是要 l = mid+1 的
// 如果 mid 小于等于 n-1 呢？
// 我们需要保证缩进到最小值
// 因此我们需要基于 mid 和 right 的比较？
// func findMin(nums []int) int {
// 	l, r := 0, len(nums)-1
// 	for l < r {
// 		mid := l + (r-l)/2
// 		if nums[mid] < nums[r] {
// 			r = mid
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return nums[r]
// }

// // 只要 mid 比最右边的小，我们就 r = mid
// // 其余情况，mid 肯定都在左侧的局部递增区间，因此可以直接 l = mid+1
// func findMin(nums []int) int {
// 	n := len(nums)
// 	l, r := 0, n-1
// 	for l < r {
// 		mid := l + (r-l)/2
// 		if nums[mid] < nums[n-1] {
// 			r = mid
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return nums[r]
// }
