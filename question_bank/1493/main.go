package main

// https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/
// 1493. 删掉一个元素以后全为 1 的最长子数组

// 不定长滑动窗口
// 窗口性质-窗口内只包含1(允许一个0)
// 这里要注意，这里是必须删掉一个元素，而不是修改元素
// 因此 len = r-l
func longestSubarray(nums []int) int {
	last := -1
	var l, r, maxLen int
	for r < len(nums) {
		if nums[r] == 0 {
			if last >= l {
				l = last + 1
			}
			last = r
		}
		maxLen = max(maxLen, r-l)
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
