package main

// https://leetcode.cn/problems/max-consecutive-ones-iii/
// 1004. 最大连续1的个数 III

// 这道题做完还可以继续做 424. 替换后的最长重复字符，思路基本一样

// 最多可以翻转 k 个 0
// 不定长滑动窗口，窗口性质，窗口内都是1(可以有k个0在窗口内)
func longestOnes(nums []int, k int) int {
	var l, r, maxLen int
	for r < len(nums) {
		// 这里算正常 r 右移扩大窗口，如果等于 1 就相当于不处理，正常扩大
		if nums[r] == 0 {
			k--
		}
		// 如果 k 小于 0，说明窗口内不满足都是1，存在超出k个0的情况
		// 此时需要不断去右移 l，并判断移除的数是否为 0，是的话 k++
		for k < 0 {
			if nums[l] == 0 {
				k++
			}
			l++
		}
		r++
		maxLen = max(maxLen, r-l)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
