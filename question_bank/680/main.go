package main

// https://leetcode.cn/problems/valid-palindrome-ii/
// 680. 验证回文串 II

// s 都由小写字符组成，能删除一个字母，即有一个容错机会
// 但是当两个字符不等时，我们可以删除左边的，也可删除右边的
// 因此两边都需要尝试，只要有一边成功就可以返回
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			// 错了之后看  [l+1,r] 和 [l,r-1] 是不是回文串
			// 是就直接返回，不是就返回错误了
			tl, tr := l+1, r
			for tl <= tr {
				if s[tl] != s[tr] {
					break
				}
				tl++
				tr--
			}
			if tl >= tr {
				return true
			}
			tl, tr = l, r-1
			for tl <= tr {
				if s[tl] != s[tr] {
					break
				}
				tl++
				tr--
			}
			if tl >= tr {
				return true
			}
			return false
		}
	}
	return true
}
