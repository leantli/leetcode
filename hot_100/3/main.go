package main

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/?favorite=2cktkvj
// 3. 无重复字符的最长子串

// 1. 最基础做法，O(n^2) 遍历，指向一个字符后，一直遍历到遇到下一个相同字符或末尾，再计算子串长度

// 2. 当然，一次遍历最好，因此我们仍需依靠 map，空间换时间
// 还需要双指针，保证能够计算长度
// 每遍历到一个字符，就将该字符存入 map，value 为当前下标
// 每遇到一个字符，就查看该字符上次出现的位置，是否大于等于左指针，
// 是的话就要移动左指针位置了，避免子串中出现重复字符

func lengthOfLongestSubstring(s string) int {
	bs := []byte(s)
	m := make(map[byte]int)
	l := 0
	var res int
	for r := 0; r < len(bs); r++ {
		if _, ok := m[bs[r]]; ok {
			if m[bs[r]] >= l {
				l = m[bs[r]] + 1
			}
		}
		diff := r - l + 1
		if res < diff {
			res = diff
		}
		m[bs[r]] = r
	}
	return res
}
