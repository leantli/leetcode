package main

// https://leetcode.cn/problems/unique-substrings-in-wraparound-string/
// 467. 环绕字符串中唯一的子字符串

// 最难的其实是去重的思路吧

// 看到了别的大佬超简洁的代码
// 这样看来就是一个很明了的滑窗了
// 窗口性质--最长递增子串s
// 满足条件后开始右移窗口左边界
// 不满足条件后开始右移窗口右边界
// func findSubstringInWraproundString(s string) int {
// 	cnt, ans := [26]int{}, 0
// 	for l, r := 0, 1; r <= len(s); r++ {
// 		if r == len(s) || (s[r]-s[r-1]+26)%26 != 1 {
// 			for ; l < r; l++ {
// 				if cnt[s[l]-'a'] < r-l {
// 					cnt[s[l]-'a'] = r - l
// 				}
// 			}
// 		}
// 	}
// 	for _, c := range cnt {
// 		ans += c
// 	}
// 	return ans
// }

// 看了题解，重新反思
// 仍是需要一个 26 长度的数组，存储每个字母开头的最长连续子串的长度
// 但这里我们并不是直接基于26个字母去进行遍历
// 而是直接遍历 s 中的每个字母，然后更新数组中各个字母开头的最长连续子串长度
// 在一定程度上减少了之前过多的重复遍历
// 而且我们判断子串递增时，也无需再额外用一个 base 字符串和 cur 下标，只需判断子串前后是否差为 1
func findSubstringInWraproundString(s string) int {
	var charLongest [26]int
	for l := 0; l < len(s); l++ {
		// 这里直接双指针找最长递增子串，无需额外一个 base 字符串和 cur 下标
		// 注意 z 和 a 这里环形递增，26取模
		r := l + 1
		for r < len(s) && s[r]%26 == (s[r-1]+1)%26 {
			r++
		}
		if charLongest[s[l]-'a'] < r-l {
			charLongest[s[l]-'a'] = r - l
		}
		// 我们求出了 l~r 这一段是最长连续递增子串了，那么[l,r]区间的所有字符都可以直接做一次判断
		// 而无需重复遍历滑窗找最长递增子串
		for l+1 != r {
			l++
			if charLongest[s[l]-'a'] < r-l {
				charLongest[s[l]-'a'] = r - l
			}
		}
	}
	var res int
	for _, lens := range charLongest {
		res += lens
	}
	return res
}

// // 窗口中的字符是连续的
// // 但主要是 非空子串 的统计，是需要去重的，在这种条件下进行这种去重感觉很麻烦
// // 比如说"abcdabcde"
// // 此时窗口肯定会检验到两个连续子串"abcd""abcde"，此时"abcd"产出的满足条件的子串数量就重复计算了
// // 如果靠排列组合+set，绝对会超时，整个时间复杂度极高
// // 因此这里可以考虑到，其实26个字母又是一个极小的数据量
// // 这里我们可以枚举26个字母开头，求各个字母开头的最长连续子串(满足环绕字符串的条件下)
// // 此时计算即可26个字母开头的最长连续子串情况下的符合条件的数目即可
// func findSubstringInWraproundString(s string) int {
// 	base := "abcdefghijklmnopqrstuvwxyz"
// 	var cnt [26]int
// 	for i := 0; i < 26; i++ {
// 		var l, r, maxLen int
// 		for r < len(s) {
// 			cur := i
// 			if s[r] != base[cur%26] {
// 				l++
// 				r++
// 				continue
// 			}
// 			for r < len(s) && s[r] == base[cur%26] {
// 				cur++
// 				r++
// 			}
// 			if maxLen < r-l {
// 				maxLen = r - l
// 			}
// 			l = r
// 		}
// 		cnt[i] = maxLen
// 	}
// 	var ans int
// 	for _, cnt := range cnt {
// 		ans += cnt
// 	}
// 	return ans
// }
