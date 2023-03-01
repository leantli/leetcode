package main

// https://leetcode.cn/problems/valid-palindrome/
// 125. 验证回文串

// 双指针，分别从头和尾部开始遍历
// 遇到新的字母再停止并比较
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		for l <= r && !isLetter(s[l]) {
			l++
		}
		for l <= r && !isLetter(s[r]) {
			r--
		}
		if l > r {
			break
		}
		if !isEqual(s[l], s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func isEqual(a, b byte) bool {
	if a > 'Z' {
		a -= 32
	}
	if b > 'Z' {
		b -= 32
	}
	return a == b
}

func isLetter(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}
