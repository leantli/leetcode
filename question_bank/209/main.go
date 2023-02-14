package main

// https://leetcode.cn/problems/minimum-size-subarray-sum/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c7qm097
// 209. 长度最小的子数组

// 在滑动窗口题又遇到了
// 不定长滑动窗口，窗口性质为 窗口内数值和大于等于 target
// 不满足条件时 r右移扩大窗口，满足条件时右移l缩小窗口
// 缩小前先判断好长度，记录最短的满足条件的窗口长度
func minSubArrayLen(target int, nums []int) int {
	var l, r, curSum int
	minLen := len(nums) + 1
	for r < len(nums) {
		curSum += nums[r]
		r++
		for curSum >= target {
			minLen = min(minLen, r-l)
			curSum -= nums[l]
			l++
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// n 个正整数， 一个 target
// 找到 一个 长度最小的 连续子数组，数组和 大于等于 target
// 返回长度，不存在返回 0

// 时间复杂度为 O(n) 和 O(nlogn)
// 最小长度显然是，单个数就已经大于等于 target
// 首先 时间复杂度为 O(n) 的话，应该是使用滑动窗口，思路比较简单？
// 反而是 nlogn，考虑一下，肯定要用二分对吧，要二分显然需要单调性，要么排序，要么前缀和！
// 这里显然无需排序，因为我们需要的是连续的子数组，排序后反而不利于判断子数组连续
// 这里首先肯定是需要去计算各个子数组的和，但是感觉不好计算，最佳方式是采用前缀和？
// 先计算出整个数组每个位置的前缀和，并且由于原数组都是正整数，因此其前缀和数组是递增的，更方便进行二分
// 再遍历每个前缀和，做二分，[l,r] 取值范围为 【当前前缀和的下标，最终前缀和的下标】
// 去寻找 右侧第一个前缀和下标，其满足 指向的前缀和 减去 当前前缀和 >= target
// 再去计算长度取最短
// func minSubArrayLen(target int, nums []int) int {
// 	// ans 为最短长度
// 	ans := math.MaxInt
// 	n := len(nums)
// 	sum := make([]int, n+1)
// 	// sum[i] 表示 前 i 个数的和-前缀和
// 	for i := 1; i <= n; i++ {
// 		sum[i] = sum[i-1] + nums[i-1]
// 	}

// 	for i := 0; i < n; i++ {
// 		// 二分，注意边界取值，这里是对 sum 的边界取值，左右都为开区间
// 		l, r := i-1, n+1
// 		for l+1 != r {
// 			mid := l + (r-l)/2
// 			if sum[mid]-sum[i] < target {
// 				l = mid
// 			} else {
// 				r = mid
// 			}
// 		}
// 		// 此时二分结束后，l 位于最后一个小于 target 的位置，r 位于第一个大于等于 target 的位置
// 		// 由于题目 target 必定大于 0，且 sum[0] 为 0，因此无需担心 l 的边界问题，只需考虑 r 可能仍在 n+1 下标位置
// 		// 当 r 仍为 边界值，则没找到大于等于 target 的前缀和，因此可以直接略过
// 		if r == n+1 {
// 			continue
// 		}
// 		ans = min(ans, r-i)
// 	}

// 	if ans == math.MaxInt {
// 		return 0
// 	}
// 	return ans
// }

// // 再优化一下滑动窗口的写法
// func minSubArrayLen(target int, nums []int) int {
// 	ans := math.MaxInt
// 	l, r, sum := 0, 0, 0
// 	for r < len(nums) {
// 		// 先模拟右扩
// 		sum += nums[r]
// 		r++
// 		// 当 sum >= target 后再缩左
// 		for sum >= target {
// 			ans = min(r-l, ans)
// 			sum -= nums[l]
// 			l++
// 		}
// 	}
// 	if ans == math.MaxInt {
// 		return 0
// 	}
// 	return ans
// }

// // 先用滑动窗口试试吧
// func minSubArrayLen(target int, nums []int) int {
// 	ans := math.MaxInt
// 	l, r, sum := 0, 0, 0
// 	for l < len(nums) {
// 		for r < len(nums) && sum < target {
// 			sum += nums[r]
// 			r++
// 		}
// 		// 出循环后有三种情况：r 已经遍历完了但是 sum 没达标； r 未遍历完但 sum 达标了，r遍历完了且 sum 刚好达标
// 		if sum < target {
// 			break
// 		}
// 		for l < r && sum >= target {
// 			ans = min(r-l, ans)
// 			sum -= nums[l]
// 			l++
// 		}
// 	}
// 	if ans == math.MaxInt {
// 		return 0
// 	}
// 	return ans
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }
