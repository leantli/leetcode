package main

// https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/
// 395. 至少有 K 个重复字符的最长子串

// 不定滑动窗口？
// 窗口性质--子串中每一字符出现次数都不少于 k
// 但是不少于 k 这个性质感觉很难用滑窗？
// 因为正常情况下窗口逐步右扩时，窗口中数量都不一定有 k，更谈何每一字符出现次数都不少于 k ？
// 也就是说，这找不到窗口滑动扩缩的条件
// 此时我们需要关注 数据范围小 的数据值 ---> 这里可以关注到 s 仅由小写英文字母组成
// 此时我们想到，不管这个子串有多长，它里面出现的字符种类，只能是[1,26]，
// 我们只需枚举子串中出现的字符种类的数量，在这个基础上，设置窗口中容纳的字符种类数量，此时窗口就能够有充足的滑动条件了
// 去滑动窗口求最长子串, 这个子串的基本条件是，容纳的字符种类数量为窗口规定的，
// 而我们需要在这个基础上，再额外检验其中各个字符出现的次数是否不少于k, 满足条件的话，我们才记录这个子串的长度，这个才是我们要的结果
// 最终得到的满足条件的最长子串，就是我们需要的结果
func longestSubstring(s string, k int) int {
	var res int
	// capacity 代表窗口中需要容纳字符的种类数量
	for capacity := 1; capacity <= 26; capacity++ {
		// 窗口中需要 capacity 个出现次数不少于 k 字符种类
		// 每当窗口中有一个字符出现次数不少于 k，则 need--
		need := capacity
		// 记录窗口中当前的字符种类数量
		var capCnt int
		var exist [26]int
		// 当窗口中字符种类不超过 capacity 时，右边界持续右扩，超过时左边界持续右扩
		var l, r int
		for r < len(s) {
			// 右扩时，增加窗口中该字符的出现次数，并且判断是否为新增的字符种类
			c := s[r] - 'a'
			exist[c]++
			if exist[c] == 1 {
				capCnt++
			}
			// 看看这个字符是否满足条件，是的话 need--
			if exist[c] == k {
				need--
			}
			for capCnt > capacity {
				lc := s[l] - 'a'
				exist[lc]--
				if exist[lc] == 0 {
					capCnt--
				}
				if exist[lc] == k-1 {
					need++
				}
				l++
			}
			r++
			// 记录最长满足条件的子串
			if need == 0 {
				if res < r-l {
					res = r - l
				}
			}
		}
	}
	return res
}
