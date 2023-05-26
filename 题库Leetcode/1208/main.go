package main

// https://leetcode.cn/problems/get-equal-substrings-within-budget/
// 1208. 尽可能使字符串相等

// 不定长窗口，窗口性质-->窗口内的 |s[i]-t[i]| 之和 <= maxCost
func equalSubstring(s string, t string, maxCost int) int {
	// 因此我们可以先计算出每个位置转换的开销
	n := len(s)
	bs, tb := []byte(s), []byte(t)
	cost := make([]int, n)
	for i, sc := range bs {
		cost[i] = abs(sc, tb[i])
	}
	// 不定长滑窗
	var l, r, maxLen, curCost int
	for r < n {
		curCost += cost[r]
		for curCost > maxCost {
			curCost -= cost[l]
			l++
		}
		r++
		maxLen = max(maxLen, r-l)
	}
	return maxLen
}

func abs(a, b byte) int {
	res := int(a) - int(b)
	if res < 0 {
		return -res
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
