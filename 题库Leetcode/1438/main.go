package main

// https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
// 1438. 绝对差不超过限制的最长连续子数组

// 再看了一下官解，最佳解答也是不定长滑窗+单调队列
// 但是和我的操作是有区别的，逻辑看起来也会更清晰简单
// 无需刻意地去比较新进数和原窗口内的最大最小值
// 我们正常的加入新进数，正常地去增加修改单调队列
// 最后判断当前单调队列的最大最小值绝对差是否超过 limit 即可
// 超过再右移 l
func longestSubarray(nums []int, limit int) int {
	var l, r int
	// 单调递增队列，首部是最小值
	// 单调递减队列，首部是最大值
	incQue, desQue := make([]int, 0), make([]int, 0)
	res := 1
	for r < len(nums) {
		// 根据新进数更新两个单调队列
		// 更新单调递增队列
		for len(incQue) > 0 && incQue[len(incQue)-1] > nums[r] {
			incQue = incQue[:len(incQue)-1]
		}
		incQue = append(incQue, nums[r])
		// 更新单调递减队列
		for len(desQue) > 0 && desQue[len(desQue)-1] < nums[r] {
			desQue = desQue[:len(desQue)-1]
		}
		desQue = append(desQue, nums[r])
		// 判断当前窗口中的最大值最小值绝对差值是否超过 limit
		// 是的话则维护窗口性质
		for abs(desQue[0], incQue[0]) > limit {
			if incQue[0] == nums[l] {
				incQue = incQue[1:]
			}
			if desQue[0] == nums[l] {
				desQue = desQue[1:]
			}
			l++
		}
		r++
		if r-l > res {
			res = r - l
		}
	}
	return res
}

func abs(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}

// // 首先这肯定要用不定长滑动窗口
// // 窗口性质---窗口内任意两数的绝对差值小于等于 limit
// // 考虑什么时候会破坏这个窗口的性质？
// // 新进数位于窗口最大值最小值之间，那肯定没问题
// // 但是如果是最大值最小值之外，可能就会导致绝对差值超过 limit
// // 因此我们可以维护窗口内的最大最小值，每次新进数时，计算该数与窗口内最大最小值的差是否超过 limit
// // 是的话就移动 l++
// // 但是这个时候会发现，仅仅维护两个最大最小值是不行的，因为 l++ 后，可能最大最小值就出去了
// // 显然这个时候我们得依靠单调队列----不理解的话见 239. 滑动窗口最大值
// func longestSubarray(nums []int, limit int) int {
// 	var l, r int
// 	// 单调递增队列，首部是最小值
// 	// 单调递减队列，首部是最大值
// 	incQue, desQue := append([]int{}, nums[r]), append([]int{}, nums[r])
// 	r++
// 	res := 1
// 	for r < len(nums) {
// 		// 新加进来一个数，判断它会不会导致窗口不满足性质---窗口内任意两数绝对差小于等于limit
// 		// 是的话就保证窗口性质，直接舍弃最左边的数看看，但是这里我们得保证 min 和 max
// 		// 所以我们得维护一个窗口中的单调递增队列和单调递减队列
// 		for (len(incQue) > 0 && abs(nums[r], incQue[0]) > limit) || (len(desQue) > 0 && abs(nums[r], desQue[0]) > limit) {
// 			if incQue[0] == nums[l] {
// 				incQue = incQue[1:]
// 			}
// 			if desQue[0] == nums[l] {
// 				desQue = desQue[1:]
// 			}
// 			l++
// 		}
// 		// 接下来更新两个单调队列
// 		// 单调递增
// 		for len(incQue) > 0 && incQue[len(incQue)-1] > nums[r] {
// 			incQue = incQue[:len(incQue)-1]
// 		}
// 		incQue = append(incQue, nums[r])
// 		// 单调递减
// 		for len(desQue) > 0 && desQue[len(desQue)-1] < nums[r] {
// 			desQue = desQue[:len(desQue)-1]
// 		}
// 		desQue = append(desQue, nums[r])
// 		r++
// 		if r-l > res {
// 			res = r - l
// 		}
// 	}
// 	return res
// }

// func abs(a, b int) int {
// 	res := a - b
// 	if res < 0 {
// 		return -res
// 	}
// 	return res
// }
