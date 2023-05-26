package main

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/?favorite=2cktkvj
// 3. 无重复字符的最长子串

// 第二次刷到
// 找出不含重复子串的最长子串
// 不定长滑动窗口？
// 窗口性质---窗口内不存在重复字符
// 如何维护？
// 维护窗口内已有字符的下标
// 当遇见重复字符时，l直接右移到该重复字符的右侧
// 但是值维护窗口内已有字符，那么每次l右移都还要考虑去除出窗口的
// 其实好像也没必要，只有当重复的字符大于等于l时，才需要移动l
// 直接维护一个每个字符上次出现的位置即可
func lengthOfLongestSubstring(s string) int {
	var l, r int
	var res int
	last := make(map[byte]int)
	for r < len(s) {
		if idx, ok := last[s[r]]; ok {
			if idx >= l {
				l = idx + 1
			}
		}
		last[s[r]] = r
		r++
		if res < r-l {
			res = r - l
		}
	}
	return res
}

// 1. 最基础做法，O(n^2) 遍历，指向一个字符后，一直遍历到遇到下一个相同字符或末尾，再计算子串长度

// // 2. 当然，一次遍历最好，因此我们仍需依靠 map，空间换时间
// // 还需要双指针，保证能够计算长度
// // 每遍历到一个字符，就将该字符存入 map，value 为当前下标
// // 每遇到一个字符，就查看该字符上次出现的位置，是否大于等于左指针，
// // 是的话就要移动左指针位置了，避免子串中出现重复字符
// func lengthOfLongestSubstring(s string) int {
// 	bs := []byte(s)
// 	m := make(map[byte]int)
// 	l := 0
// 	var res int
// 	for r := 0; r < len(bs); r++ {
// 		if _, ok := m[bs[r]]; ok {
// 			if m[bs[r]] >= l {
// 				l = m[bs[r]] + 1
// 			}
// 		}
// 		diff := r - l + 1
// 		if res < diff {
// 			res = diff
// 		}
// 		m[bs[r]] = r
// 	}
// 	return res
// }
