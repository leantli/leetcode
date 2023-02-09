package main

// https://leetcode.cn/problems/minimum-window-substring/
// 76. 最小覆盖子串

// 刚才那种滑窗，每次都完全地比较了子串，并且直接记录字符串，每次变更时开销较大
// 优化了完全比较两个字符串的判断：基于子串 有什么字符，以及这些字符出现多少次
// 我们只需更注重窗口内-那些子串有的字符，看这些字符是否大于等于子串中该字符出现的次数
// 是的话，该字符就满足了包含子串的条件，当所有子串有的字符，都满足条件时，就说明当前窗口包含子串t
// 比如说 "abcdefg" 和 "abc"
// 我们只需要关注前者中的'a”b”c'三个字符，判断他们出现的次数是否大于等于后者中该三个字符出现的次数
func minWindow(s string, t string) string {
	n, m := len(s), len(t)
	// 特殊处理，t 长度超了s 直接返回
	if m > n {
		return ""
	}
	// 初始化子串各个字符的出现次数以及有几个字符需要满足出现次数以上
	tMap, need := make(map[byte]int), 0
	for _, c := range []byte(t) {
		if tMap[c] == 0 {
			need++
		}
		tMap[c]++
	}
	// 初始化滑动窗口，以及窗口内已满足包含子串的字符数量
	// 当 satisfy == need 时，即可将对应的子串进行记录
	var l, r, satisfy, minL int
	windowMap := make([]int, 128)
	bs := []byte(s)
	minLen := n + 1
	for r < n {
		windowMap[bs[r]]++
		// 这里要注意，我们需要的是 bs[r] 为 tmap 中的字符，并且出现次数大于 tmap 中该字符出现的次数，satisfy 才能增长
		if count, ok := tMap[bs[r]]; ok {
			if windowMap[bs[r]] == count {
				satisfy++
			}
		}
		r++
		// 当满足条件时，先记录对应的子串起始下标及其长度，再逐渐右移 l 缩减窗口
		for satisfy == need {
			if r-l < minLen {
				minLen = r - l
				minL = l
			}
			// 当移除的字符为 子串包含字符时，再查看该字符出现次数是否太少了，太少了则 减少满足的字符
			windowMap[bs[l]]--
			if count, ok := tMap[bs[l]]; ok {
				if windowMap[bs[l]] < count {
					satisfy--
				}
			}
			l++
		}
	}
	if minLen == n+1 {
		return ""
	}
	return string(bs[minL : minL+minLen])
}

// // 俺自己写的，不过其实这边验证是否为子串的时候还是有不少重复且浪费的验证
// // 最小子串 该子串中包含 t 中所有字符，包括重复的字符
// // 先写个确认t是否为s的子串的函数？
// // 维护一个不定长滑动窗口，统计窗口内各个字符的出现次数
// // 当满足 t 为 s 子串时，l 右移缩小窗口，否则 r 右移扩大窗口
// // 同时记录最小子串
// func minWindow(s string, t string) string {
// 	n, m := len(s), len(t)
// 	if m > n {
// 		return ""
// 	}
// 	tMap := make(map[byte]int)
// 	for _, c := range []byte(t) {
// 		tMap[c]++
// 	}
// 	windowsMap := make(map[byte]int)
// 	var l, r int
// 	bs := []byte(s)
// 	res := make([]byte, n+2)
// 	// t不为窗口的子串时 r右移
// 	// t为窗口子串时，判断此时窗口的长度，如果比此前记录的短，则将其进行记录；l右移
// 	for r < n {
// 		windowsMap[bs[r]]++
// 		for aSubB(windowsMap, tMap) {
// 			if r-l+1 < len(res) {
// 				res = bs[l : r+1]
// 			}
// 			windowsMap[bs[l]]--
// 			l++
// 		}
// 		r++
// 	}
// 	if len(res) == n+2 {
// 		return ""
// 	}
// 	return string(res)
// }

// // b是a的子串
// func aSubB(a, b map[byte]int) bool {
// 	// 当 b 中某个字符的出现次数比 a 中该字符出现的次数多，则说明 b 不为 a 的子串
// 	for k, v := range b {
// 		if v > a[k] {
// 			return false
// 		}
// 	}
// 	return true
// }
