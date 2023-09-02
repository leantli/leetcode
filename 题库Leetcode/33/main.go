package main

// https://leetcode.cn/problems/search-in-rotated-sorted-array/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 33. 搜索旋转排序数组

// 二刷
// 时间复杂度要求 logn，基本大概率使用 二分操作
// 二分就是排除掉不可能的区域
func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {

			return mid
		}
		// 这里根据 mid 去先区分好哪边是有序的
		// 比如 nums[mid] > nums[n-1]，那么此时 nums[0:mid] 是有序的，此时方便操作
		if nums[mid] > nums[n-1] {
			if target > nums[n-1] && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target <= nums[n-1] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// // 这里先考虑一下，假设是 0,1,2,4,5,6,7
// // 旋转后变成了         4,5,6,7,0,1,2
// // 要求 logn 大概率得用二分解决了
// // 这里怎么使用二分呢？
// // 我们要保证，每次缩减时，往正确的方向去缩减
// // 先写一下看看情况----这里是找准确值，因此采用 l <= r，并且 l 和 r 的 mid 更新一定都会加减，因为已经确认了 mid 不是 target
// func search(nums []int, target int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return -1
// 	}
// 	if n == 1 {
// 		if nums[0] == target {
// 			return 0
// 		} else {
// 			return -1
// 		}
// 	}
// 	l, r := 0, n-1
// 	for l <= r {
// 		mid := l + (r-l)/2
// 		if nums[mid] == target {
// 			return mid
// 		}
// 		// 接下来就是最重要的地方，
// 		// 需要保证下面的判断，使得每次 l 和 r 缩减的方向是正确的
// 		// if nums[mid] > target ?
// 		// 首先绝对不能直接基于 target 去判断，因为现在整个 nums 并不是全局升序的，而是局部升序的
// 		// 我们需要先确定 mid 处于哪个升序区间，才能进一步向正确的方向去二分
// 		// 那么怎么判断现在 mid 位于哪个升序区间？
// 		// 最佳方法就是 mid 与 最后一个数比较，因为最后一个数是分隔局部升序的一个标志点，当然，第一个数也可以，同理的
// 		// 写了一会发现，单纯知道 mid 在哪个升序区间其实还是不行的，我们还需要知道 target 应该位于哪个升序区间
// 		// 否则我们在缩减时，并不一定会往正确的方向去缩减

// 		// 这里先根据 mid 的位置进行区分
// 		if nums[mid] > nums[n-1] {
// 			// 根据 nums[mid] > nums[n-1]，我们只能 保证 [0, mid] 这部分，必定是升序有序的
// 			// 因此，我们只能对 target > nums[n-1] && nums[mid] > target 的情况做处理
// 			// 此时 r = mid - 1
// 			// 其他情况，不管是 target <= nums[n-1] 还是 nums[mid] < target，都只能 l = mid + 1
// 			if target > nums[n-1] && nums[mid] > target {
// 				r = mid - 1
// 			} else {
// 				l = mid + 1
// 			}
// 		} else {
// 			// 同理，只能保证 [mid, n-1] 部分必然是升序的，因此只能对如下条件做处理
// 			if target <= nums[n-1] && nums[mid] < target {
// 				l = mid + 1
// 			} else {
// 				r = mid - 1
// 			}
// 		}
// 	}
// 	return -1
// }
