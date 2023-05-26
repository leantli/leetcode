package main

// https://leetcode.cn/problems/minimum-swaps-to-group-all-1s-together-ii/
// 2134. 最少交换次数来组合所有的 1 II

// 思路和 1151. 最少交换次数来组合所有的 1 一样
// 只是多了环形数组的小处理

func minSwaps(nums []int) int {
	var oneCount int
	for _, num := range nums {
		if num == 1 {
			oneCount++
		}
	}
	// 维护一个 oneCount 长度的窗口，我们知道必定 oneCount <= len(nums)
	var oneCnt int // 用于计算窗口中已有的 1 的数量
	// 最终只要用 oneCount 减去窗口最大的 1 数量，就是最小交换次数(交换0)
	for _, num := range nums[:oneCount] {
		if num == 1 {
			oneCnt++
		}
	}
	maxOneCnt := oneCnt
	n := len(nums)
	for i := oneCount; i < n+oneCount; i++ {
		r, l := nums[i%n], nums[(i-oneCount)%n]
		oneCnt += r - l
		if maxOneCnt < oneCnt {
			maxOneCnt = oneCnt
		}
	}
	return oneCount - maxOneCnt
}
