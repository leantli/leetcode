package main

// https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 48. 最长不含重复字符的子字符串

// 看到后马上想到一个方法，算是双指针？
// 一次遍历配合 map：遇到一个字母，就看看map里面有没有，没有就是第一次出现
// 有的话就看看这个字母上次出现的位置是否大于 l 左指针
// 大于的话就移动 l 指针到上次出现的位置处 +1
// 统一存储当前字母及其位置，顺便计算下子串长度
func lengthOfLongestSubstring(s string) int {
	bs := []byte(s)
	n := len(bs)
	m := make(map[byte]int)
	res := 0
	l := 0
	for i := 0; i < n; i++ {
		if _, ok := m[bs[i]]; ok {
			if m[bs[i]] >= l {
				l = m[bs[i]] + 1
			}
		}
		m[bs[i]] = i
		res = Max(res, i-l+1)
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
