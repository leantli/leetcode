package main

// https://leetcode.cn/problems/minimum-window-subsequence/
// 727. 最小窗口子序列

// 在前者基础上的一点小优化
// 之前就是满足条件后逆序找一次，只右移了l，但是只有再遇到t的最后一个字符，就又要回溯，可能存在过多的重复情况
// 因此这次，找到符合条件的之后，l右移至该子串的开头，r也要移到l的右侧，重新开始新的条件遍历，能减少更多重复的遍历判断
func minWindow(s string, t string) string {
	if s == "" || t == "" || len(s) < len(t) {
		return ""
	}
	tIndex := 0
	var l, r int
	resL, resR := 0, 20001
	for r < len(s) {
		if s[r] == t[tIndex] {
			tIndex++
		}
		if tIndex == len(t) {
			end := r
			for end >= l {
				if s[end] == t[tIndex-1] {
					tIndex--
				}
				if tIndex == 0 {
					if r-end+1 < resR-resL {
						resL = end
						resR = r + 1
					}
					r = end + 1
					break
				}
				end--
			}
			l = end
		}
		r++
	}
	if resR == 20001 {
		return ""
	}
	return string(s[resL:resR])
}

// // 在 s 中找到一个最短的子串 w，使得 t 是 w 的子序列
// // 不定长滑窗，窗口性质- t 是窗口内的字符串的子序列
// // 当窗口满足性质时，(记录当前最短长度的字符串+缩左侧)直至不满足窗口性质，
// // 不满足后再继续右扩
// // 这道题要做应该是不难的，但是每次都完全地判断 t 是否为窗口内字符串的子序列，开销其实过大
// // 因此我们得想想其他方法
// // 什么时候才开始往窗口塞东西 ? t[index] 表示窗口下一个要匹配的 t 的字符
// // 当 s[i] = t[0] 时开始往窗口塞东西
// // 但是什么时候窗口的左边界右移呢？
// // 首先肯定是 s[i] = t[len(t)-1] 时，将 s[i] 加入到窗口之后，此时窗口满足性质
// // 此时在窗口中逆序遍历，依次找到 t[len(t)-1], t[len(t)-2] ... t[0]
// // 此时找到 t[0] 后，在窗口内的逆序遍历就可以了停止了，此时逆序的长度即当前窗口的最短连续子串长度
// // 此时窗口的左边界也移动至这个 t[0] 的位置，之后又可以继续去右移 r
// // 并且只要遇到 t 最后一个字符，就可以继续逆序找
// func minWindow(s string, t string) string {
// 	if s == "" || t == "" || len(s) < len(t) {
// 		return ""
// 	}
// 	tIndex := 0
// 	var l, r int
// 	resL, resR := 0, 20001
// 	for r < len(s) {
// 		if s[r] == t[tIndex] {
// 			tIndex++
// 		}
// 		if tIndex == len(t) {
// 			end := r
// 			for end >= l {
// 				if s[end] == t[tIndex-1] {
// 					tIndex--
// 				}
// 				if tIndex == 0 {
// 					if r-end+1 < resR-resL {
// 						resL = end
// 						resR = r + 1
// 					}
// 					tIndex = len(t) - 1
// 					break
// 				}
// 				end--
// 			}
// 			l = end
// 		}
// 		r++
// 	}
// 	if resR == 20001 {
// 		return ""
// 	}
// 	return string(s[resL:resR])
// }
