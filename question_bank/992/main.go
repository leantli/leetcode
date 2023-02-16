package main

// https://leetcode.cn/problems/subarrays-with-k-different-integers/
// 992. K 个不同整数的子数组

// 不定长滑动窗口
// 窗口性质--窗口中不同整数的个数恰好为 k
// 恰好显然是达不到的，我们可以转换为至多 k 个的好子数组的数目
// 接下来再一次滑窗求 至多 k-1 个的好子数组的数目
// 此时前者-后者即为 恰好 k 个的好子数组的数目
func subarraysWithKDistinct(nums []int, k int) int {
	count := func(k int) int {
		var res int
		// 窗口右扩时判断给对应的下标++
		// 当这个下标对应的 val 为 1 时，说明刚刚从 0 到 1
		// 此时窗口内的不同整数数量+1
		// 超过 k 时左边界右移
		var l, r, curCnt int
		exist := make([]int, len(nums)+1)
		for r < len(nums) {
			exist[nums[r]]++
			if exist[nums[r]] == 1 {
				curCnt++
			}
			// 如果超了，不满足窗口要求，则维护窗口性质
			for curCnt > k {
				exist[nums[l]]--
				if exist[nums[l]] == 0 {
					curCnt--
				}
				l++
			}
			r++
			res += r - l
		}
		return res
	}
	return count(k) - count(k-1)
}
