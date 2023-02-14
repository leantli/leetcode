package main

// https://leetcode.cn/problems/max-consecutive-ones-ii/
// 487. 最大连续1的个数 II

func findMaxConsecutiveOnes(nums []int) int {
	// 因为最多只能翻转一个 0，所以我们用一个值 表示上次出现 0 的下标
	lastZero := -1
	var l, r, maxLen int
	for r < len(nums) {
		// 遇到 0 时，保证窗口内性质即可，窗口内性质为，全部都为 1 (有一个位置可以为 0)
		// 如果遇到 0，就判断上个 0 是否已经在窗口内，是的话
		// 就移动窗口的左边界到上个 0 的右侧，并且更新 last 为当前遇到的 0
		if nums[r] == 0 {
			if lastZero >= l {
				l = lastZero + 1
			}
			lastZero = r
		}
		maxLen = max(maxLen, r-l+1)
		r++
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
