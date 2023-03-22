package main

// https://leetcode.cn/problems/isomorphic-strings/
// 205. 同构字符串

// 其需要一个映射，并且 s 和 t 都是由 ascii 组成
// 没有关联的两个字母就先关联上，两边都要维护各自的一个表
// 避免同个字母被另一边多个新字符映射
// 如果有关联的字母，发现对不上，就返回false
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm, tm := make([]byte, 128), make([]byte, 128)
	for i := range s {
		if sm[s[i]] == 0 && tm[t[i]] == 0 {
			sm[s[i]] = t[i]
			tm[t[i]] = s[i]
			continue
		}
		if sm[s[i]] == t[i] && tm[t[i]] == s[i] {
			continue
		}
		return false
	}
	return true
}

// // 这个网上看到的也不错，如果两个字符相映射
// // 那么这两个字符在各自的字符串中，第一次出现的位置一定是一样的
// // 不过感觉没有上面的容易想到
// func isIsomorphic(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	for i := range s {
// 		if strings.IndexByte(s, s[i]) != strings.IndexByte(t, t[i]) {
// 			return false
// 		}
// 	}
// 	return true
// }
