package main

// 题目：给你一个整数数组 nums 和两个整数：left 及 right 。找出 nums 中连续、非空且其中最大元素在范围 [left, right] 内的子数组，并返回满足条件的子数组的个数
// https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum/
// 795. 区间子数组个数

// 滑窗题再相遇
// 连续，窗口中最大元素在 [left,right]范围内
// 窗口性质，窗口内没有元素大于 right + 窗口内需要有元素大于 left
// 返回满足条件的子数组个数
// 这里其实是可以问题转换，但是感觉还会挺难想到的
// 分别求 窗口内没有元素大于 right 和 窗口内没有元素大于 left 的子数组数量
// 此时 前者 - 后者，得到的就是窗口内最大元素在 [left,right] 范围内
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	count := func(threshold int) int {
		var l, r, cnt int
		for r < len(nums) {
			// 维护窗口性质
			if nums[r] > threshold {
				l = r + 1
			}
			r++
			cnt += r - l
		}
		return cnt
	}
	return count(right) - count(left-1)
}

// // 只会暴力遍历，超时，想了挺久没想到解法，看了官方题解发现第一种也有点难想到
// // 第二种问题转换比较好考虑
// // 要求找到 nums 中连续、非空且其中最大元素在范围 [left, right] 内的子数组
// // 我们可以考虑将其转换为 [0,..right] 然后再减去 [0,..left−1] 的问题
// // 也就是说，所有最大元素不超过 right 的子数组个数，减去所有最大元素不超过 left-1 的子数组个数
// // 剩下的就是最大元素在区间 [left,..right] 范围内的子数组个数，即题目要求的结果
// // 问题转换，两区见计数
// func numSubarrayBoundedMax(nums []int, left int, right int) int {
// 	count := func(num int) int {
// 		t := 0
// 		cnt := 0
// 		for _, v := range nums {
// 			if v > num {
// 				t = 0
// 				continue
// 			}
// 			t++
// 			cnt += t
// 		}
// 		return cnt
// 	}
// 	return count(right) - count(left-1)
// }

// 官方解第一种方法
// 满足题目的要求有如下两个约束：
// 1. 子数组中至少有一个元素满足 >= left；
// 2. 子数组中没有元素满足 > right
// 为了满足第一个条件，我们只需要从左向右遍历数组，直到遇到第一个 >= left 的元素，将 r = i；
// 同样地，为了满足第二个条件，我们需要找到第一个 > right 的元素，将 l = i；
// 由此分为以下几种情况：
// 当 nums[i] > right 时，此时必然也会 nums[i] > left，所以区间是 0；
// 当 nums[i] >= left && nums[i] < right 时，此时区间是 [l + 1, r]；
// 当 nums[i] < left 时，加和上一个区间的范围，即如果是情况一加和 0，情况二加和区间范围；
// 好有技巧性
// func numSubarrayBoundedMax(nums []int, left int, right int) (res int) {
// 	l, r := -1, -1
// 	for i, x := range nums {
// 		if x > right {
// 			l = i
// 		}
// 		if x >= left {
// 			r = i
// 		}
// 		res += r - l
// 	}
// 	return
// }

// // 先写一个判断元素是否在 [left, right] 内的函数
// // 暴力遍历-->超时
// func numSubarrayBoundedMax(nums []int, left int, right int) int {
// 	res := 0
// 	for l := 0; l < len(nums); l++ {
// 		max := nums[l]
// 		for r := l; r < len(nums); r++ {
// 			if max < nums[r] {
// 				max = nums[r]
// 			}
// 			if !numIsVaild(max, left, right) {
// 				continue
// 			}
// 			res++
// 		}
// 	}
// 	return res
// }

// func numIsVaild(num, left, right int) bool {
// 	if num < left || num > right {
// 		return false
// 	}
// 	return true
// }
