package main

// https://leetcode.cn/problems/reverse-vowels-of-a-string/
// 345. 反转字符串中的元音字母

// 双指针从首位开始
// 分别向中间遍历，遇到元音字母则停止，两边都停止时交换元音字母的位置
// 再继续向中间遍历
func reverseVowels(s string) string {
	bs := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isVowels(bs[l]) {
			l++
		}
		for l < r && !isVowels(bs[r]) {
			r--
		}
		bs[l], bs[r] = bs[r], bs[l]
		l++
		r--
	}
	return string(bs)
}

func isVowels(c byte) bool {
	if c <= 'Z' {
		c += 32
	}
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}
	return false
}
